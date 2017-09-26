// Package graphrbac implements the Azure ARM Graphrbac service API version 1.6.
//
// The Graph RBAC Management Client
package graphrbac

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.22.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Graphrbac
	DefaultBaseURI = "https://graph.windows.net"
)

// ManagementClient is the base client for Graphrbac.
type ManagementClient struct {
	autorest.Client
	BaseURI  string
	TenantID string
}

// New creates an instance of the ManagementClient client.
func New(tenantID string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, tenantID)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, tenantID string) ManagementClient {
	return ManagementClient{
		Client:   autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:  baseURI,
		TenantID: tenantID,
	}
}
