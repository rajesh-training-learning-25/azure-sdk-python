//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azfile

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type connection struct {
	u string
	p runtime.Pipeline
}

// newConnection creates an instance of the connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func newConnection(endpoint string, options *azcore.ClientOptions) *connection {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	return &connection{u: endpoint, p: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, &cp)}
}

// Endpoint returns the connection's endpoint.
func (c *connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *connection) Pipeline() runtime.Pipeline {
	return c.p
}
