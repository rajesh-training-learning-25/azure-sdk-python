// +build go1.9

// Copyright 2017 Microsoft Corporation
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
// commit ID: 09b7f75d0b75dab01d5152a968422915e4cb76ce

package resources

import original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-02-01/resources"

type DeploymentOperationsClient = original.DeploymentOperationsClient
type GroupsClient = original.GroupsClient
type DeploymentMode = original.DeploymentMode

const (
	Complete	DeploymentMode	= original.Complete
	Incremental	DeploymentMode	= original.Incremental
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type AliasPathType = original.AliasPathType
type AliasType = original.AliasType
type BasicDependency = original.BasicDependency
type DebugSetting = original.DebugSetting
type Dependency = original.Dependency
type Deployment = original.Deployment
type DeploymentExportResult = original.DeploymentExportResult
type DeploymentExtended = original.DeploymentExtended
type DeploymentExtendedFilter = original.DeploymentExtendedFilter
type DeploymentListResult = original.DeploymentListResult
type DeploymentOperation = original.DeploymentOperation
type DeploymentOperationProperties = original.DeploymentOperationProperties
type DeploymentOperationsListResult = original.DeploymentOperationsListResult
type DeploymentProperties = original.DeploymentProperties
type DeploymentPropertiesExtended = original.DeploymentPropertiesExtended
type DeploymentValidateResult = original.DeploymentValidateResult
type ExportTemplateRequest = original.ExportTemplateRequest
type GenericResource = original.GenericResource
type GenericResourceFilter = original.GenericResourceFilter
type Group = original.Group
type GroupExportResult = original.GroupExportResult
type GroupFilter = original.GroupFilter
type GroupListResult = original.GroupListResult
type GroupProperties = original.GroupProperties
type HTTPMessage = original.HTTPMessage
type Identity = original.Identity
type ListResult = original.ListResult
type ManagementErrorWithDetails = original.ManagementErrorWithDetails
type MoveInfo = original.MoveInfo
type ParametersLink = original.ParametersLink
type Plan = original.Plan
type Provider = original.Provider
type ProviderListResult = original.ProviderListResult
type ProviderOperationDisplayProperties = original.ProviderOperationDisplayProperties
type ProviderResourceType = original.ProviderResourceType
type Resource = original.Resource
type Sku = original.Sku
type SubResource = original.SubResource
type TagCount = original.TagCount
type TagDetails = original.TagDetails
type TagsListResult = original.TagsListResult
type TagValue = original.TagValue
type TargetResource = original.TargetResource
type TemplateLink = original.TemplateLink

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type DeploymentsClient = original.DeploymentsClient
type GroupClient = original.GroupClient
type ProvidersClient = original.ProvidersClient
type TagsClient = original.TagsClient

func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClient(subscriptionID)
}
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewProvidersClient(subscriptionID string) ProvidersClient {
	return original.NewProvidersClient(subscriptionID)
}
func NewProvidersClientWithBaseURI(baseURI string, subscriptionID string) ProvidersClient {
	return original.NewProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsClient(subscriptionID string) TagsClient {
	return original.NewTagsClient(subscriptionID)
}
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return original.NewTagsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewGroupsClient(subscriptionID string) GroupsClient {
	return original.NewGroupsClient(subscriptionID)
}
func NewGroupsClientWithBaseURI(baseURI string, subscriptionID string) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewDeploymentOperationsClient(subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClient(subscriptionID)
}
func NewDeploymentOperationsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClientWithBaseURI(baseURI, subscriptionID)
}
