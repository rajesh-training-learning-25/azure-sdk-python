// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package admin

import (
	"context"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/atom"
)

// Client allows you to administer resources in a Service Bus Namespace.
// For example, you can create queues, enabling capabilities like partitioning, duplicate detection, etc..
// NOTE: For sending and receiving messages you'll need to use the `azservicebus.Client` type instead.
type Client struct {
	em atom.EntityManager
}

type ClientOptions struct {
	// for future expansion
}

// NewClientFromConnectionString creates a Client authenticating using a connection string.
func NewClientFromConnectionString(connectionString string, options *ClientOptions) (*Client, error) {
	em, err := atom.NewEntityManagerWithConnectionString(connectionString, internal.Version)

	if err != nil {
		return nil, err
	}

	return &Client{em: em}, nil
}

// NewClient creates a Client authenticating using a TokenCredential.
func NewClient(fullyQualifiedNamespace string, tokenCredential azcore.TokenCredential, options *ClientOptions) (*Client, error) {
	em, err := atom.NewEntityManager(fullyQualifiedNamespace, tokenCredential, internal.Version)

	if err != nil {
		return nil, err
	}

	return &Client{em: em}, nil
}

type NamespaceProperties struct {
	CreatedTime  time.Time
	ModifiedTime time.Time

	SKU            string
	MessagingUnits *int64
	Name           string
}

type GetNamespacePropertiesResult struct {
	NamespaceProperties
}

type GetNamespacePropertiesResponse struct {
	GetNamespacePropertiesResult
	RawResponse *http.Response
}

type GetNamespacePropertiesOptions struct {
	// For future expansion
}

// GetNamespaceProperties gets the properties for the namespace, includings properties like SKU and CreatedTime.
func (ac *Client) GetNamespaceProperties(ctx context.Context, options *GetNamespacePropertiesOptions) (*GetNamespacePropertiesResponse, error) {
	var body *atom.NamespaceEntry
	resp, err := ac.em.Get(ctx, "/$namespaceinfo", &body)

	if err != nil {
		return nil, err
	}

	return &GetNamespacePropertiesResponse{
		RawResponse: resp,
		GetNamespacePropertiesResult: GetNamespacePropertiesResult{
			NamespaceProperties: NamespaceProperties{
				Name:           body.NamespaceInfo.Name,
				CreatedTime:    body.NamespaceInfo.CreatedTime,
				ModifiedTime:   body.NamespaceInfo.ModifiedTime,
				SKU:            body.NamespaceInfo.MessagingSKU,
				MessagingUnits: body.NamespaceInfo.MessagingUnits,
			},
		},
	}, nil
}
