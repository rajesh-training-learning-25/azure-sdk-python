// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/url"
	"strings"
)

const scope = "https://management.azure.com//.default"
const telemetryInfo = "azsdk-go-armstorage/<version>"

// ClientOptions contains configuration settings for the default client's pipeline.
type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
	// ApplicationID is an application-specific identification string used in telemetry.
	// It has a maximum length of 24 characters and must not contain any spaces.
	ApplicationID string
}

// DefaultClientOptions creates a ClientOptions type initialized with default values.
func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

func (c *ClientOptions) telemetryOptions() azcore.TelemetryOptions {
	t := telemetryInfo
	if c.ApplicationID != "" {
		a := strings.ReplaceAll(c.ApplicationID, " ", "/")
		if len(a) > 24 {
			a = a[:24]
		}
		t = fmt.Sprintf("%s %s", a, telemetryInfo)
	}
	if c.Telemetry.Value == "" {
		return azcore.TelemetryOptions{Value: t}
	}
	return azcore.TelemetryOptions{Value: fmt.Sprintf("%s %s", c.Telemetry.Value, t)}
}

// Client - The Azure Storage Management API.
type Client struct {
	u *url.URL
	p azcore.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "https://management.azure.com"

// NewDefaultClient creates an instance of the Client type using the DefaultEndpoint.
func NewDefaultClient(cred azcore.Credential, options *ClientOptions) (*Client, error) {
	return NewClient(DefaultEndpoint, cred, options)
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, cred azcore.Credential, options *ClientOptions) (*Client, error) {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}}),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		return nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)
	}
	return &Client{u: u, p: p}, nil
}

// Operations returns the Operations associated with this client.
func (client *Client) Operations() Operations {
	return &operations{Client: client}
}

// SkusOperations returns the SkusOperations associated with this client.
func (client *Client) SkusOperations(subscriptionID string) SkusOperations {
	return &skusOperations{Client: client, subscriptionID: subscriptionID}
}

// StorageAccountsOperations returns the StorageAccountsOperations associated with this client.
func (client *Client) StorageAccountsOperations(subscriptionID string) StorageAccountsOperations {
	return &storageAccountsOperations{Client: client, subscriptionID: subscriptionID}
}

// UsagesOperations returns the UsagesOperations associated with this client.
func (client *Client) UsagesOperations(subscriptionID string) UsagesOperations {
	return &usagesOperations{Client: client, subscriptionID: subscriptionID}
}

// ManagementPoliciesOperations returns the ManagementPoliciesOperations associated with this client.
func (client *Client) ManagementPoliciesOperations(subscriptionID string) ManagementPoliciesOperations {
	return &managementPoliciesOperations{Client: client, subscriptionID: subscriptionID}
}

// PrivateEndpointConnectionsOperations returns the PrivateEndpointConnectionsOperations associated with this client.
func (client *Client) PrivateEndpointConnectionsOperations(subscriptionID string) PrivateEndpointConnectionsOperations {
	return &privateEndpointConnectionsOperations{Client: client, subscriptionID: subscriptionID}
}

// PrivateLinkResourcesOperations returns the PrivateLinkResourcesOperations associated with this client.
func (client *Client) PrivateLinkResourcesOperations(subscriptionID string) PrivateLinkResourcesOperations {
	return &privateLinkResourcesOperations{Client: client, subscriptionID: subscriptionID}
}

// ObjectReplicationPoliciesOperations returns the ObjectReplicationPoliciesOperations associated with this client.
func (client *Client) ObjectReplicationPoliciesOperations(subscriptionID string) ObjectReplicationPoliciesOperations {
	return &objectReplicationPoliciesOperations{Client: client, subscriptionID: subscriptionID}
}

// EncryptionScopesOperations returns the EncryptionScopesOperations associated with this client.
func (client *Client) EncryptionScopesOperations(subscriptionID string) EncryptionScopesOperations {
	return &encryptionScopesOperations{Client: client, subscriptionID: subscriptionID}
}

// BlobServicesOperations returns the BlobServicesOperations associated with this client.
func (client *Client) BlobServicesOperations(subscriptionID string) BlobServicesOperations {
	return &blobServicesOperations{Client: client, subscriptionID: subscriptionID}
}

// BlobContainersOperations returns the BlobContainersOperations associated with this client.
func (client *Client) BlobContainersOperations(subscriptionID string) BlobContainersOperations {
	return &blobContainersOperations{Client: client, subscriptionID: subscriptionID}
}

// FileServicesOperations returns the FileServicesOperations associated with this client.
func (client *Client) FileServicesOperations(subscriptionID string) FileServicesOperations {
	return &fileServicesOperations{Client: client, subscriptionID: subscriptionID}
}

// FileSharesOperations returns the FileSharesOperations associated with this client.
func (client *Client) FileSharesOperations(subscriptionID string) FileSharesOperations {
	return &fileSharesOperations{Client: client, subscriptionID: subscriptionID}
}
