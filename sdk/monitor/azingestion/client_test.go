//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azingestion_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azlog "github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azingestion"
)

func TestUpload(t *testing.T) {
	azlog.SetListener(func(cls azlog.Event, msg string) {
		fmt.Println(msg)
	})
	endpoint := os.Getenv("MONITOR_INGESTION_DATA_COLLECTION_ENDPOINT")
	ruleID := os.Getenv("INGESTION_DATA_COLLECTION_RULE_IMMUTABLE_ID")
	streamName := os.Getenv("INGESTION_STREAM_NAME")
	// clientID := os.Getenv("AZINGESTION_CLIENT_ID")
	// clientSecret := os.Getenv("AZINGESTION_CLIENT_SECRET")
	// tenantID := os.Getenv("AZINGESTION_TENANT_ID")

	// credential, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	client, err := azingestion.NewClient(endpoint, credential, &azingestion.ClientOptions{azcore.ClientOptions{Logging: policy.LogOptions{IncludeBody: true}}})
	if err != nil {
		panic(err)
	}

	data := "[{\"Time\":\"2023-04-24T13:17:33.8008175\",\"Computer\":\"Computer1\",\"AdditionalContext\":2},{\"Time\":\"2023-04-24T13:17:33.8008175\",\"Computer\":\"Computer2\",\"AdditionalContext\":3}]}"

	res, err := client.Upload(context.Background(), ruleID, streamName, data, nil)
	if err != nil {
		panic(err)
	}
	_ = res

}
