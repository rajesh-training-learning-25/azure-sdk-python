// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azlogs_test

import (
	"context"
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/query/azlogs"
	"github.com/stretchr/testify/require"
)

const (
	recordingDirectory  = "sdk/monitor/query/azlogs/testdata"
	fakeWorkspaceID     = "32d1e136-gg81-4b0a-b647-260cdc471f68"
	fakeWorkspaceID2    = "asdjkfj8k20-gg81-4b0a-9fu2-260c09fn1f68"
	fakeResourceURI     = "/subscriptions/faa080af-c1d8-40ad-9cce-e1a451va7b87/resourceGroups/rg-example/providers/Microsoft.AppConfiguration/configurationStores/example"
	fakeSubscrtiptionID = "faa080af-c1d8-40ad-9cce-e1a451va7b87"
	fakeRegion          = "westus2"
)

var (
	credential     azcore.TokenCredential
	workspaceID    string
	workspaceID2   string
	resourceURI    string
	subscriptionID string
	region         string
	clientCloud    cloud.Configuration
)

func TestMain(m *testing.M) {
	code := run(m)
	os.Exit(code)
}

func run(m *testing.M) int {
	if recording.GetRecordMode() == recording.PlaybackMode || recording.GetRecordMode() == recording.RecordingMode {
		proxy, err := recording.StartTestProxy(recordingDirectory, nil)
		if err != nil {
			panic(err)
		}
		defer func() {
			err := recording.StopTestProxy(proxy)
			if err != nil {
				panic(err)
			}
		}()
	}

	if recording.GetRecordMode() == recording.PlaybackMode {
		credential = &FakeCredential{}
	} else {
		tenantID := getEnvVar("AZLOGS_TENANT_ID", "")
		clientID := getEnvVar("AZLOGS_CLIENT_ID", "")
		secret := getEnvVar("AZLOGS_CLIENT_SECRET", "")
		var err error
		credential, err = azidentity.NewClientSecretCredential(tenantID, clientID, secret, nil)
		if err != nil {
			panic(err)
		}

		if cloudEnv, ok := os.LookupEnv("AZLOGS_ENVIRONMENT"); ok {
			if strings.EqualFold(cloudEnv, "AzureUSGovernment") {
				clientCloud = cloud.AzureGovernment
			}
			if strings.EqualFold(cloudEnv, "AzureChinaCloud") {
				clientCloud = cloud.AzureChina
			}
		}

	}
	workspaceID = getEnvVar("WORKSPACE_ID", fakeWorkspaceID)
	workspaceID2 = getEnvVar("WORKSPACE_ID2", fakeWorkspaceID2)
	resourceURI = getEnvVar("RESOURCE_URI", fakeResourceURI)
	subscriptionID = getEnvVar("AZLOGS_SUBSCRIPTION_ID", fakeSubscrtiptionID)
	region = getEnvVar("AZLOGS_LOCATION", fakeRegion)

	return m.Run()
}

func startRecording(t *testing.T) {
	err := recording.Start(t, recordingDirectory, nil)
	require.NoError(t, err)
	t.Cleanup(func() {
		err := recording.Stop(t, nil)
		require.NoError(t, err)
	})
}

func startTest(t *testing.T) *azlogs.Client {
	startRecording(t)
	transport, err := recording.NewRecordingHTTPClient(t, nil)
	require.NoError(t, err)
	opts := &azlogs.ClientOptions{ClientOptions: azcore.ClientOptions{Transport: transport, Cloud: clientCloud}}
	client, err := azlogs.NewClient(credential, opts)
	if err != nil {
		panic(err)
	}
	return client
}

func getEnvVar(envVar string, fakeValue string) string {
	// get value
	value := fakeValue
	if recording.GetRecordMode() == recording.LiveMode || recording.GetRecordMode() == recording.RecordingMode {
		value = os.Getenv(envVar)
		if value == "" {
			panic("no value for " + envVar)
		}
	}

	// sanitize value
	if fakeValue != "" && recording.GetRecordMode() == recording.RecordingMode {
		err := recording.AddGeneralRegexSanitizer(fakeValue, value, nil)
		if err != nil {
			panic(err)
		}
	}

	return value
}

type FakeCredential struct{}

func (f *FakeCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{Token: "faketoken", ExpiresOn: time.Now().Add(time.Hour).UTC()}, nil
}

type serdeModel interface {
	json.Marshaler
	json.Unmarshaler
}

func testSerde[T serdeModel](t *testing.T, model T) {
	data, err := model.MarshalJSON()
	require.NoError(t, err)
	err = model.UnmarshalJSON(data)
	require.NoError(t, err)

	// testing unmarshal error scenarios
	var data2 []byte
	err = model.UnmarshalJSON(data2)
	require.Error(t, err)

	m := regexp.MustCompile(":.*$")
	modifiedData := m.ReplaceAllString(string(data), ":false}")
	if !strings.Contains(modifiedData, "render") && modifiedData != "{}" {
		data3 := []byte(modifiedData)
		err = model.UnmarshalJSON(data3)
		require.Error(t, err)
	}
}
