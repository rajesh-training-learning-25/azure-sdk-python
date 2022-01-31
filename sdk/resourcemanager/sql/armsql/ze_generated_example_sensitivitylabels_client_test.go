//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsListByDatabaseWithSourceCurrent.json
func ExampleSensitivityLabelsClient_ListCurrentByDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	pager := client.ListCurrentByDatabase("<resource-group-name>",
		"<server-name>",
		"<database-name>",
		&armsql.SensitivityLabelsClientListCurrentByDatabaseOptions{SkipToken: nil,
			Count:  nil,
			Filter: nil,
		})
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

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsCurrentUpdate.json
func ExampleSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.SensitivityLabelUpdateList{
			Operations: []*armsql.SensitivityLabelUpdate{
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsql.SensitivityLabelUpdateKindSet.ToPtr(),
						SensitivityLabel: &armsql.SensitivityLabel{
							Properties: &armsql.SensitivityLabelProperties{
								InformationType:   to.StringPtr("<information-type>"),
								InformationTypeID: to.StringPtr("<information-type-id>"),
								LabelID:           to.StringPtr("<label-id>"),
								LabelName:         to.StringPtr("<label-name>"),
								Rank:              armsql.SensitivityLabelRankLow.ToPtr(),
							},
						},
						Table: to.StringPtr("<table>"),
					},
				},
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsql.SensitivityLabelUpdateKindSet.ToPtr(),
						SensitivityLabel: &armsql.SensitivityLabel{
							Properties: &armsql.SensitivityLabelProperties{
								InformationType:   to.StringPtr("<information-type>"),
								InformationTypeID: to.StringPtr("<information-type-id>"),
								LabelID:           to.StringPtr("<label-id>"),
								LabelName:         to.StringPtr("<label-name>"),
								Rank:              armsql.SensitivityLabelRankCritical.ToPtr(),
							},
						},
						Table: to.StringPtr("<table>"),
					},
				},
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsql.SensitivityLabelUpdateKindRemove.ToPtr(),
						Table:  to.StringPtr("<table>"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsListByDatabaseWithSourceRecommended.json
func ExampleSensitivityLabelsClient_ListRecommendedByDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	pager := client.ListRecommendedByDatabase("<resource-group-name>",
		"<server-name>",
		"<database-name>",
		&armsql.SensitivityLabelsClientListRecommendedByDatabaseOptions{SkipToken: nil,
			IncludeDisabledRecommendations: nil,
			Filter:                         nil,
		})
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

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/RecommendedColumnSensitivityLabelEnable.json
func ExampleSensitivityLabelsClient_EnableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	_, err = client.EnableRecommendation(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		"<schema-name>",
		"<table-name>",
		"<column-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/RecommendedColumnSensitivityLabelDisable.json
func ExampleSensitivityLabelsClient_DisableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	_, err = client.DisableRecommendation(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		"<schema-name>",
		"<table-name>",
		"<column-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnSensitivityLabelGet.json
func ExampleSensitivityLabelsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		"<schema-name>",
		"<table-name>",
		"<column-name>",
		armsql.SensitivityLabelSourceCurrent,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SensitivityLabelsClientGetResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnSensitivityLabelCreateMax.json
func ExampleSensitivityLabelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		"<schema-name>",
		"<table-name>",
		"<column-name>",
		armsql.SensitivityLabel{
			Properties: &armsql.SensitivityLabelProperties{
				InformationType:   to.StringPtr("<information-type>"),
				InformationTypeID: to.StringPtr("<information-type-id>"),
				LabelID:           to.StringPtr("<label-id>"),
				LabelName:         to.StringPtr("<label-name>"),
				Rank:              armsql.SensitivityLabelRankLow.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SensitivityLabelsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnSensitivityLabelDelete.json
func ExampleSensitivityLabelsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		"<schema-name>",
		"<table-name>",
		"<column-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
