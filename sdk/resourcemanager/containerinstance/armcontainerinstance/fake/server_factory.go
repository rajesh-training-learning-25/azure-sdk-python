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

// ServerFactory is a fake server for instances of the armcontainerinstance.ClientFactory type.
type ServerFactory struct {
	ContainerGroupsServer              ContainerGroupsServer
	ContainersServer                   ContainersServer
	LocationServer                     LocationServer
	OperationsServer                   OperationsServer
	SubnetServiceAssociationLinkServer SubnetServiceAssociationLinkServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armcontainerinstance.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armcontainerinstance.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                  *ServerFactory
	trMu                                 sync.Mutex
	trContainerGroupsServer              *ContainerGroupsServerTransport
	trContainersServer                   *ContainersServerTransport
	trLocationServer                     *LocationServerTransport
	trOperationsServer                   *OperationsServerTransport
	trSubnetServiceAssociationLinkServer *SubnetServiceAssociationLinkServerTransport
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
	case "ContainerGroupsClient":
		initServer(s, &s.trContainerGroupsServer, func() *ContainerGroupsServerTransport {
			return NewContainerGroupsServerTransport(&s.srv.ContainerGroupsServer)
		})
		resp, err = s.trContainerGroupsServer.Do(req)
	case "ContainersClient":
		initServer(s, &s.trContainersServer, func() *ContainersServerTransport { return NewContainersServerTransport(&s.srv.ContainersServer) })
		resp, err = s.trContainersServer.Do(req)
	case "LocationClient":
		initServer(s, &s.trLocationServer, func() *LocationServerTransport { return NewLocationServerTransport(&s.srv.LocationServer) })
		resp, err = s.trLocationServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SubnetServiceAssociationLinkClient":
		initServer(s, &s.trSubnetServiceAssociationLinkServer, func() *SubnetServiceAssociationLinkServerTransport {
			return NewSubnetServiceAssociationLinkServerTransport(&s.srv.SubnetServiceAssociationLinkServer)
		})
		resp, err = s.trSubnetServiceAssociationLinkServer.Do(req)
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
