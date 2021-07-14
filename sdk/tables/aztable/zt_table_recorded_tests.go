// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztable

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/assert"
)

type tablesRecordedTests struct{}

type testContext struct {
	recording *recording.Recording
	client    *TableServiceClient
	context   *recording.TestContext
}

const (
	storageAccountNameEnvVar    = "TABLES_STORAGE_ACCOUNT_NAME"
	cosmosAccountNameEnnVar     = "TABLES_COSMOS_ACCOUNT_NAME"
	storageEndpointSuffixEnvVar = "STORAGE_ENDPOINT_SUFFIX"
	cosmosEndpointSuffixEnvVar  = "COSMOS_TABLES_ENDPOINT_SUFFIX"
	storageAccountKeyEnvVar     = "TABLES_PRIMARY_STORAGE_ACCOUNT_KEY"
	cosmosAccountKeyEnvVar      = "TABLES_PRIMARY_COSMOS_ACCOUNT_KEY"
	tableNamePrefix             = "gotable"
	DefaultStorageSuffix        = "core.windows.net"
	DefaultCosmosSuffix         = "cosmos.azure.com"
)

type EndpointType string

const (
	StorageEndpoint EndpointType = "storage"
	CosmosEndpoint  EndpointType = "cosmos"
)

var ctx = context.Background()
var clientsMap map[string]*testContext = make(map[string]*testContext)

func storageURI(accountName string, endpointSuffix string) string {
	return "https://" + accountName + ".table." + endpointSuffix
}

func cosmosURI(accountName string, endpointSuffix string) string {
	return "https://" + accountName + ".table." + endpointSuffix
}

// create the test specific TableClient and wire it up to recordings
func recordedTestSetup(t *testing.T, testName string, endpointType EndpointType, mode recording.RecordMode) {
	var accountName string
	var suffix string
	var cred *SharedKeyCredential
	var secret string
	var uri string
	assert := assert.New(t)

	// init the test framework
	context := recording.NewTestContext(func(msg string) { assert.FailNow(msg) }, func(msg string) { t.Log(msg) }, func() string { return testName })
	r, err := recording.NewRecording(context, mode)
	assert.Nil(err)

	if endpointType == StorageEndpoint {
		accountName, err = r.GetRecordedVariable(storageAccountNameEnvVar, recording.Default)
		suffix = r.GetOptionalRecordedVariable(storageEndpointSuffixEnvVar, DefaultStorageSuffix, recording.Default)
		secret, err = r.GetRecordedVariable(storageAccountKeyEnvVar, recording.Secret_Base64String)
		cred, _ = NewSharedKeyCredential(accountName, secret)
		uri = storageURI(accountName, suffix)
	} else {
		accountName, err = r.GetRecordedVariable(cosmosAccountNameEnnVar, recording.Default)
		suffix = r.GetOptionalRecordedVariable(cosmosEndpointSuffixEnvVar, DefaultCosmosSuffix, recording.Default)
		secret, err = r.GetRecordedVariable(cosmosAccountKeyEnvVar, recording.Secret_Base64String)
		cred, _ = NewSharedKeyCredential(accountName, secret)
		uri = cosmosURI(accountName, suffix)
	}

	client, err := NewTableServiceClient(uri, cred, &TableClientOptions{HTTPClient: r, Retry: azcore.RetryOptions{MaxRetries: -1}})

	assert.Nil(err)
	clientsMap[testName] = &testContext{client: client, recording: r, context: &context}
}

func recordedTestTeardown(key string) {
	context, ok := clientsMap[key]
	if ok && !(*context.context).IsFailed() {
		context.recording.Stop()
	}
}

// cleans up the specified tables. If tables is nil, all tables will be deleted
func cleanupTables(context *testContext, tables *[]string) {
	c := context.client
	if tables == nil {
		pager := c.Query(QueryOptions{})
		for pager.NextPage(ctx) {
			for _, t := range pager.PageResponse().TableQueryResponse.Value {
				c.Delete(ctx, *t.TableName)
			}
		}
	} else {
		for _, t := range *tables {
			c.Delete(ctx, t)
		}
	}
}

func getTestContext(key string) *testContext {
	return clientsMap[key]
}

func getTableName(context *testContext, prefix ...string) (string, error) {
	if len(prefix) == 0 {
		return context.recording.GenerateAlphaNumericID(tableNamePrefix, 20, true)
	} else {
		return context.recording.GenerateAlphaNumericID(prefix[0], 20, true)
	}
}

type basicTestEntity struct {
	Entity
	Integer int32
	String  string
	Bool    bool
}

type complexTestEntity struct {
	Entity
	Integer  int
	String   string
	Bool     bool
	Float    float32
	DateTime time.Time
	Byte     []byte
	// Integer64 int64 // Need to add type hints for ints/floats above 32bits
	// Float64   float64
}

func createSimpleEntity(count int, pk string) basicTestEntity {
	return basicTestEntity{
		Entity: Entity{
			PartitionKey: pk,
			RowKey:       fmt.Sprint(count),
		},
		String:  fmt.Sprintf("some string %d", count),
		Integer: int32(count),
		Bool:    true,
	}
}

func createSimpleEntities(count int, pk string) *[]basicTestEntity {
	result := make([]basicTestEntity, count)
	for i := 1; i <= count; i++ {
		result[i-1] = createSimpleEntity(i, pk)
	}
	return &result
}

// func createComplexMapEntities(context *testContext, count int, pk string) *[]map[string]interface{} {
// 	result := make([]map[string]interface{}, count)

// 	for i := 1; i <= count; i++ {
// 		var e = map[string]interface{}{
// 			partitionKey:          pk,
// 			rowKey:                fmt.Sprint(i),
// 			"StringProp":          fmt.Sprintf("some string %d", i),
// 			"IntProp":             i,
// 			"BoolProp":            true,
// 			"SomeBinaryProperty":  []byte("some bytes"),
// 			"SomeDateProperty":    context.recording.Now(),
// 			"SomeDoubleProperty0": float64(1),
// 			"SomeDoubleProperty1": float64(1.2345),
// 			"SomeGuidProperty":    context.recording.UUID(),
// 			"SomeInt64Property":   (int64)(math.MaxInt64),
// 			"SomeIntProperty":     42,
// 			"SomeStringProperty":  "some string",
// 		}
// 		result[i-1] = e
// 	}
// 	return &result
// }

func createComplexEntities(context *testContext, count int, pk string) *[]complexTestEntity {
	result := make([]complexTestEntity, count)

	for i := 1; i <= count; i++ {
		var e = complexTestEntity{
			Entity: Entity{
				PartitionKey: "partition",
				RowKey:       "row",
			},
			Integer:  int(i),
			String:   "someString",
			Bool:     true,
			Float:    3.14159,
			DateTime: time.Now(),
			Byte:     []byte("somebytes"),
			// Integer64: int64(math.Pow(2, 33)),
			// Float64:   math.Pow(2, 33.1),
		}
		result[i-1] = e
	}
	return &result
}

type simpleEntity struct {
	ETag         string
	PartitionKey string
	RowKey       string
	Timestamp    time.Time
	IntProp      int
	BoolProp     bool
	StringProp   string
}

type complexEntity struct {
	ETag                  string
	PartitionKey          string
	RowKey                string
	Timestamp             time.Time
	SomeBinaryProperty    []byte
	SomeDateProperty      time.Time
	SomeDoubleProperty0   float64
	SomeDoubleProperty1   float64
	SomeGuidProperty      [16]byte `uuid:""`
	SomeInt64Property     int64
	SomeIntProperty       int
	SomeStringProperty    string
	SomePtrStringProperty *string
}
