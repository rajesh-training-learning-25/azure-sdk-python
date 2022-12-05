//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/stretchr/testify/suite"
)

type TableresourcesTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	accountName       string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *TableresourcesTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.accountName = "databaseaccounttable"
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/cosmos/armcosmos/testdata")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *TableresourcesTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestTableresourcesTestSuite(t *testing.T) {
	suite.Run(t, new(TableresourcesTestSuite))
}

func (testsuite *TableresourcesTestSuite) Prepare() {
	var err error
	// From step DatabaseAccount_Create
	databaseAccountsClient, err := armcosmos.NewDatabaseAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	databaseAccountsClientCreateOrUpdateResponsePoller, err := databaseAccountsClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Kind:     to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			Capabilities: []*armcosmos.Capability{
				{
					Name: to.Ptr("EnableTable"),
				}},
			CreateMode:               to.Ptr(armcosmos.CreateModeDefault),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("eastus"),
				}},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
}

// Microsoft.DocumentDB/databaseAccounts/tables
func (testsuite *TableresourcesTestSuite) TestTableresources() {
	tableName := "cosmos_table"
	var err error
	// From step TableResources_CreateTable
	tableResourcesClient, err := armcosmos.NewTableResourcesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	tableResourcesClientCreateUpdateTableResponsePoller, err := tableResourcesClient.BeginCreateUpdateTable(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, tableName, armcosmos.TableCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.TableCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.TableResource{
				ID: to.Ptr(tableName),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, tableResourcesClientCreateUpdateTableResponsePoller)
	testsuite.Require().NoError(err)

	// From step TableResources_GetTable
	_, err = tableResourcesClient.GetTable(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, tableName, nil)
	testsuite.Require().NoError(err)

	// From step TableResources_ListTables
	tableResourcesClientNewListTablesPager := tableResourcesClient.NewListTablesPager(testsuite.resourceGroupName, testsuite.accountName, nil)
	for tableResourcesClientNewListTablesPager.More() {
		_, err := tableResourcesClientNewListTablesPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step TableResources_UpdateTableThroughput
	tableResourcesClientUpdateTableThroughputResponsePoller, err := tableResourcesClient.BeginUpdateTableThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, tableName, armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, tableResourcesClientUpdateTableThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step TableResources_GetTableThroughput
	_, err = tableResourcesClient.GetTableThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, tableName, nil)
	testsuite.Require().NoError(err)

	// From step TableResources_MigrateTableToAutoscale
	tableResourcesClientMigrateTableToAutoscaleResponsePoller, err := tableResourcesClient.BeginMigrateTableToAutoscale(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, tableName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, tableResourcesClientMigrateTableToAutoscaleResponsePoller)
	testsuite.Require().NoError(err)

	// From step TableResources_MigrateTableToManualThroughput
	tableResourcesClientMigrateTableToManualThroughputResponsePoller, err := tableResourcesClient.BeginMigrateTableToManualThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, tableName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, tableResourcesClientMigrateTableToManualThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step TableResources_DeleteTable
	tableResourcesClientDeleteTableResponsePoller, err := tableResourcesClient.BeginDeleteTable(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, tableName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, tableResourcesClientDeleteTableResponsePoller)
	testsuite.Require().NoError(err)
}