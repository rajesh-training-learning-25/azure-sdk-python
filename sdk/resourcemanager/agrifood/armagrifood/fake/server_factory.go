//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armagrifood.ClientFactory type.
type ServerFactory struct {
	ExtensionsServer                 ExtensionsServer
	FarmBeatsExtensionsServer        FarmBeatsExtensionsServer
	FarmBeatsModelsServer            FarmBeatsModelsServer
	LocationsServer                  LocationsServer
	OperationsServer                 OperationsServer
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer       PrivateLinkResourcesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armagrifood.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armagrifood.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trExtensionsServer                 *ExtensionsServerTransport
	trFarmBeatsExtensionsServer        *FarmBeatsExtensionsServerTransport
	trFarmBeatsModelsServer            *FarmBeatsModelsServerTransport
	trLocationsServer                  *LocationsServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer       *PrivateLinkResourcesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "ExtensionsClient":
		initServer(s, &s.trExtensionsServer, func() *ExtensionsServerTransport { return NewExtensionsServerTransport(&s.srv.ExtensionsServer) })
		resp, err = s.trExtensionsServer.Do(req)
	case "FarmBeatsExtensionsClient":
		initServer(s, &s.trFarmBeatsExtensionsServer, func() *FarmBeatsExtensionsServerTransport {
			return NewFarmBeatsExtensionsServerTransport(&s.srv.FarmBeatsExtensionsServer)
		})
		resp, err = s.trFarmBeatsExtensionsServer.Do(req)
	case "FarmBeatsModelsClient":
		initServer(s, &s.trFarmBeatsModelsServer, func() *FarmBeatsModelsServerTransport {
			return NewFarmBeatsModelsServerTransport(&s.srv.FarmBeatsModelsServer)
		})
		resp, err = s.trFarmBeatsModelsServer.Do(req)
	case "LocationsClient":
		initServer(s, &s.trLocationsServer, func() *LocationsServerTransport { return NewLocationsServerTransport(&s.srv.LocationsServer) })
		resp, err = s.trLocationsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
