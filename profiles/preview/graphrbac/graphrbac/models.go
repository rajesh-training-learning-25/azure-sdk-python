// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package graphrbac

import original "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"

type ApplicationsClient = original.ApplicationsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DomainsClient = original.DomainsClient
type GroupsClient = original.GroupsClient
type ObjectType = original.ObjectType

const (
	ObjectTypeApplication      ObjectType = original.ObjectTypeApplication
	ObjectTypeDirectoryObject  ObjectType = original.ObjectTypeDirectoryObject
	ObjectTypeGroup            ObjectType = original.ObjectTypeGroup
	ObjectTypeServicePrincipal ObjectType = original.ObjectTypeServicePrincipal
	ObjectTypeUser             ObjectType = original.ObjectTypeUser
)

type UserType = original.UserType

const (
	Guest  UserType = original.Guest
	Member UserType = original.Member
)

type AADObject = original.AADObject
type ADGroup = original.ADGroup
type Application = original.Application
type ApplicationAddOwnerParameters = original.ApplicationAddOwnerParameters
type ApplicationCreateParameters = original.ApplicationCreateParameters
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationUpdateParameters = original.ApplicationUpdateParameters
type AppRole = original.AppRole
type CheckGroupMembershipParameters = original.CheckGroupMembershipParameters
type CheckGroupMembershipResult = original.CheckGroupMembershipResult
type BasicDirectoryObject = original.BasicDirectoryObject
type DirectoryObject = original.DirectoryObject
type DirectoryObjectListResult = original.DirectoryObjectListResult
type Domain = original.Domain
type DomainListResult = original.DomainListResult
type ErrorMessage = original.ErrorMessage
type GetObjectsParameters = original.GetObjectsParameters
type GetObjectsResult = original.GetObjectsResult
type GetObjectsResultIterator = original.GetObjectsResultIterator
type GetObjectsResultPage = original.GetObjectsResultPage
type GraphError = original.GraphError
type GroupAddMemberParameters = original.GroupAddMemberParameters
type GroupCreateParameters = original.GroupCreateParameters
type GroupGetMemberGroupsParameters = original.GroupGetMemberGroupsParameters
type GroupGetMemberGroupsResult = original.GroupGetMemberGroupsResult
type GroupListResult = original.GroupListResult
type GroupListResultIterator = original.GroupListResultIterator
type GroupListResultPage = original.GroupListResultPage
type KeyCredential = original.KeyCredential
type KeyCredentialListResult = original.KeyCredentialListResult
type KeyCredentialsUpdateParameters = original.KeyCredentialsUpdateParameters
type OdataError = original.OdataError
type PasswordCredential = original.PasswordCredential
type PasswordCredentialListResult = original.PasswordCredentialListResult
type PasswordCredentialsUpdateParameters = original.PasswordCredentialsUpdateParameters
type PasswordProfile = original.PasswordProfile
type RequiredResourceAccess = original.RequiredResourceAccess
type ResourceAccess = original.ResourceAccess
type ServicePrincipal = original.ServicePrincipal
type ServicePrincipalCreateParameters = original.ServicePrincipalCreateParameters
type ServicePrincipalListResult = original.ServicePrincipalListResult
type ServicePrincipalListResultIterator = original.ServicePrincipalListResultIterator
type ServicePrincipalListResultPage = original.ServicePrincipalListResultPage
type SignInName = original.SignInName
type User = original.User
type UserBase = original.UserBase
type UserCreateParameters = original.UserCreateParameters
type UserGetMemberGroupsParameters = original.UserGetMemberGroupsParameters
type UserGetMemberGroupsResult = original.UserGetMemberGroupsResult
type UserListResult = original.UserListResult
type UserListResultIterator = original.UserListResultIterator
type UserListResultPage = original.UserListResultPage
type UserUpdateParameters = original.UserUpdateParameters
type ObjectsClient = original.ObjectsClient
type ServicePrincipalsClient = original.ServicePrincipalsClient
type UsersClient = original.UsersClient

func NewApplicationsClient(tenantID string) ApplicationsClient {
	return original.NewApplicationsClient(tenantID)
}
func NewApplicationsClientWithBaseURI(baseURI string, tenantID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, tenantID)
}
func New(tenantID string) BaseClient {
	return original.New(tenantID)
}
func NewWithBaseURI(baseURI string, tenantID string) BaseClient {
	return original.NewWithBaseURI(baseURI, tenantID)
}
func NewDomainsClient(tenantID string) DomainsClient {
	return original.NewDomainsClient(tenantID)
}
func NewDomainsClientWithBaseURI(baseURI string, tenantID string) DomainsClient {
	return original.NewDomainsClientWithBaseURI(baseURI, tenantID)
}
func NewGroupsClient(tenantID string) GroupsClient {
	return original.NewGroupsClient(tenantID)
}
func NewGroupsClientWithBaseURI(baseURI string, tenantID string) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, tenantID)
}
func PossibleObjectTypeValues() []ObjectType {
	return original.PossibleObjectTypeValues()
}
func PossibleUserTypeValues() []UserType {
	return original.PossibleUserTypeValues()
}
func NewObjectsClient(tenantID string) ObjectsClient {
	return original.NewObjectsClient(tenantID)
}
func NewObjectsClientWithBaseURI(baseURI string, tenantID string) ObjectsClient {
	return original.NewObjectsClientWithBaseURI(baseURI, tenantID)
}
func NewServicePrincipalsClient(tenantID string) ServicePrincipalsClient {
	return original.NewServicePrincipalsClient(tenantID)
}
func NewServicePrincipalsClientWithBaseURI(baseURI string, tenantID string) ServicePrincipalsClient {
	return original.NewServicePrincipalsClientWithBaseURI(baseURI, tenantID)
}
func NewUsersClient(tenantID string) UsersClient {
	return original.NewUsersClient(tenantID)
}
func NewUsersClientWithBaseURI(baseURI string, tenantID string) UsersClient {
	return original.NewUsersClientWithBaseURI(baseURI, tenantID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
