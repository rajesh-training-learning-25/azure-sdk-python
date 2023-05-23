//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/exported"
)

func (client *AppendBlobClient) Endpoint() string {
	return client.endpoint
}

func (client *AppendBlobClient) Pipeline() runtime.Pipeline {
	return client.internal.Pipeline()
}

// NewAppendBlobClient creates a new instance of AppendBlobClient with the specified values.
//   - endpoint - The URL of the service account, container, or blob that is the target of the desired operation.
func NewAppendBlobClient(endpoint string, pl runtime.PipelineOptions, clientOptions *azcore.ClientOptions) (*AppendBlobClient, error) {
	azClient, err := azcore.NewClient("appendblob.Client", exported.ModuleVersion, pl, clientOptions)
	if err != nil {
		return nil, err
	}
	client := &AppendBlobClient{
		endpoint: endpoint,
		internal: azClient,
	}
	return client, nil
}
