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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiOperation.json
func ExampleAPIOperationClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIOperationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.GetEntityTag(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiOperation.json
func ExampleAPIOperationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIOperationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiOperation.json
func ExampleAPIOperationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIOperationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		armapimanagement.OperationContract{
			Properties: &armapimanagement.OperationContractProperties{
				Description:        to.Ptr("<description>"),
				TemplateParameters: []*armapimanagement.ParameterContract{},
				Request: &armapimanagement.RequestContract{
					Description:     to.Ptr("<description>"),
					Headers:         []*armapimanagement.ParameterContract{},
					QueryParameters: []*armapimanagement.ParameterContract{},
					Representations: []*armapimanagement.RepresentationContract{
						{
							ContentType: to.Ptr("<content-type>"),
							SchemaID:    to.Ptr("<schema-id>"),
							TypeName:    to.Ptr("<type-name>"),
						}},
				},
				Responses: []*armapimanagement.ResponseContract{
					{
						Description: to.Ptr("<description>"),
						Headers:     []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{
							{
								ContentType: to.Ptr("<content-type>"),
							},
							{
								ContentType: to.Ptr("<content-type>"),
							}},
						StatusCode: to.Ptr[int32](200),
					}},
				Method:      to.Ptr("<method>"),
				DisplayName: to.Ptr("<display-name>"),
				URLTemplate: to.Ptr("<urltemplate>"),
			},
		},
		&armapimanagement.APIOperationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApiOperation.json
func ExampleAPIOperationClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIOperationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		"<if-match>",
		armapimanagement.OperationUpdateContract{
			Properties: &armapimanagement.OperationUpdateContractProperties{
				TemplateParameters: []*armapimanagement.ParameterContract{},
				Request: &armapimanagement.RequestContract{
					QueryParameters: []*armapimanagement.ParameterContract{
						{
							Name:         to.Ptr("<name>"),
							Type:         to.Ptr("<type>"),
							Description:  to.Ptr("<description>"),
							DefaultValue: to.Ptr("<default-value>"),
							Required:     to.Ptr(true),
							Values: []*string{
								to.Ptr("sample")},
						}},
				},
				Responses: []*armapimanagement.ResponseContract{
					{
						Description:     to.Ptr("<description>"),
						Headers:         []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{},
						StatusCode:      to.Ptr[int32](200),
					},
					{
						Description:     to.Ptr("<description>"),
						Headers:         []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{},
						StatusCode:      to.Ptr[int32](500),
					}},
				Method:      to.Ptr("<method>"),
				DisplayName: to.Ptr("<display-name>"),
				URLTemplate: to.Ptr("<urltemplate>"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiOperation.json
func ExampleAPIOperationClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIOperationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		"<if-match>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
