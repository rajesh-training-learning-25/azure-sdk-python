// Deprecated: Please note, this package has been deprecated. A replacement package is available github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos. We strongly encourage you to upgrade to continue receiving updates. See Migration Guide for guidance on upgrading. Refer to our deprecation policy for more details.
//
// Package documentdb implements the Azure ARM Documentdb service API version 2021-01-15.
//
// Azure Cosmos DB Database Service Resource Provider REST API
package documentdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Documentdb
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Documentdb.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}
