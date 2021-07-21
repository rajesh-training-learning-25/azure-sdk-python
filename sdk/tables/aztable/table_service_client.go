// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztable

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const (
	LegacyCosmosTableDomain = ".table.cosmosdb."
	CosmosTableDomain       = ".table.cosmos."
)

// A TableServiceClient represents a client to the table service. It can be used to query the available tables, add/remove tables, and various other service level operations.
type TableServiceClient struct {
	client  *tableClient
	service *serviceClient
	cred    SharedKeyCredential
}

// NewTableServiceClient creates a TableServiceClient struct using the specified serviceURL, credential, and options.
func NewTableServiceClient(serviceURL string, cred azcore.Credential, options *TableClientOptions) (*TableServiceClient, error) {
	conOptions := options.getConnectionOptions()
	if isCosmosEndpoint(serviceURL) {
		conOptions.PerCallPolicies = []azcore.Policy{CosmosPatchTransformPolicy{}}
	}
	conOptions.PerCallPolicies = append(conOptions.PerCallPolicies, cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{"none"}}}))
	con := newConnection(serviceURL, conOptions)
	c, _ := cred.(*SharedKeyCredential)
	return &TableServiceClient{client: &tableClient{con}, service: &serviceClient{con}, cred: *c}, nil
}

// NewTableClient returns a pointer to a TableClient affinitzed to the specified table name and initialized with the same serviceURL and credentials as this TableServiceClient
func (t *TableServiceClient) NewTableClient(tableName string) *TableClient {
	return &TableClient{client: t.client, cred: t.cred, Name: tableName, service: t}
}

// Create creates a table with the specified name.
func (t *TableServiceClient) Create(ctx context.Context, name string) (TableResponseResponse, error) {
	resp, err := t.client.Create(ctx, TableProperties{&name}, new(TableCreateOptions), new(QueryOptions))
	if err == nil {
		tableResp := resp.(TableResponseResponse)
		return tableResp, nil
	}
	return TableResponseResponse{}, err
}

// Delete deletes a table by name.
func (t *TableServiceClient) Delete(ctx context.Context, name string) (TableDeleteResponse, error) {
	return t.client.Delete(ctx, name, nil)
}

// Query queries the existing tables using the specified QueryOptions.
// QueryOptions can specify the following properties to affect the query results returned:
//
// Filter: An Odata filter expression that limits results to those tables that satisfy the filter expression.
// For example, the following expression would return only tables with a TableName of 'foo': "TableName eq 'foo'"
//
// Top: The maximum number of tables that will be returned per page of results.
// Note: This value does not limit the total number of results if NextPage is called on the returned Pager until it returns false.
//
// Query returns a Pager, which allows iteration through each page of results. Example:
//
// pager := client.Query(nil)
// for pager.NextPage(ctx) {
//     resp = pager.PageResponse()
//     fmt.sprintf("The page contains %i results", len(resp.TableQueryResponse.Value))
// }
// err := pager.Err()
func (t *TableServiceClient) Query(queryOptions *QueryOptions) TableQueryResponsePager {
	return &tableQueryResponsePager{client: t.client, queryOptions: queryOptions, tableQueryOptions: new(TableQueryOptions)}
}

func (t *TableServiceClient) GetStatistics(ctx context.Context, options *ServiceGetStatisticsOptions) (TableServiceStatsResponse, error) {
	// TODO: This only uses the secondary endpoint which needs to be updated in the generated call
	if options == nil {
		options = &ServiceGetStatisticsOptions{}
	}
	return t.service.GetStatistics(ctx, options)
}

func (t *TableServiceClient) GetProperties(ctx context.Context, options *ServiceGetPropertiesOptions) (TableServicePropertiesResponse, error) {
	if options == nil {
		options = &ServiceGetPropertiesOptions{}
	}
	return t.service.GetProperties(ctx, options)
}

func (t *TableServiceClient) SetProperties(ctx context.Context, properties TableServiceProperties, options *ServiceSetPropertiesOptions) (ServiceSetPropertiesResponse, error) {
	if options == nil {
		options = &ServiceSetPropertiesOptions{}
	}
	return t.service.SetProperties(ctx, properties, options)
}

func isCosmosEndpoint(url string) bool {
	isCosmosEmulator := strings.Contains(url, "localhost") && strings.Contains(url, "8902")
	return isCosmosEmulator || strings.Contains(url, CosmosTableDomain) || strings.Contains(url, LegacyCosmosTableDomain)
}
