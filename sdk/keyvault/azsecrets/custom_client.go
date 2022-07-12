//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azsecrets

// this file contains handwritten additions to the generated code

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/internal"
)

// NewClient creates a client that accesses a Key Vault's secrets.
func NewClient(vaultURL string, credential azcore.TokenCredential, options *azcore.ClientOptions) *Client {
	authPolicy := internal.NewKeyVaultChallengePolicy(credential)
	pl := runtime.NewPipeline(moduleName, version, runtime.PipelineOptions{PerRetry: []policy.Policy{authPolicy}}, options)
	return &Client{endpoint: vaultURL, pl: pl}
}

// ID is a secret's unique ID, containing its name and version.
type ID string

// Name of the secret.
func (i *ID) Name() string {
	_, name, _ := internal.ParseID((*string)(i))
	return *name
}

// Version of the secret. This returns an empty string when the ID contains no version.
func (i *ID) Version() string {
	_, _, version := internal.ParseID((*string)(i))
	if version == nil {
		return ""
	}
	return *version
}
