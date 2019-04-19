// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package azurestack

import original "github.com/Azure/azure-sdk-for-go/services/azurestack/mgmt/2017-06-01/azurestack"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type CustomerSubscriptionsClient = original.CustomerSubscriptionsClient
type ComputeRole = original.ComputeRole

const (
	IaaS ComputeRole = original.IaaS
	None ComputeRole = original.None
	PaaS ComputeRole = original.PaaS
)

type OperatingSystem = original.OperatingSystem

const (
	OperatingSystemLinux   OperatingSystem = original.OperatingSystemLinux
	OperatingSystemNone    OperatingSystem = original.OperatingSystemNone
	OperatingSystemWindows OperatingSystem = original.OperatingSystemWindows
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type ActivationKeyResult = original.ActivationKeyResult
type CustomerSubscription = original.CustomerSubscription
type CustomerSubscriptionList = original.CustomerSubscriptionList
type CustomerSubscriptionListIterator = original.CustomerSubscriptionListIterator
type CustomerSubscriptionListPage = original.CustomerSubscriptionListPage
type CustomerSubscriptionProperties = original.CustomerSubscriptionProperties
type DataDiskImage = original.DataDiskImage
type Display = original.Display
type ExtendedProduct = original.ExtendedProduct
type ExtendedProductProperties = original.ExtendedProductProperties
type IconUris = original.IconUris
type Operation = original.Operation
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OsDiskImage = original.OsDiskImage
type Product = original.Product
type ProductLink = original.ProductLink
type ProductList = original.ProductList
type ProductListIterator = original.ProductListIterator
type ProductListPage = original.ProductListPage
type ProductNestedProperties = original.ProductNestedProperties
type ProductProperties = original.ProductProperties
type Registration = original.Registration
type RegistrationList = original.RegistrationList
type RegistrationListIterator = original.RegistrationListIterator
type RegistrationListPage = original.RegistrationListPage
type RegistrationParameter = original.RegistrationParameter
type RegistrationParameterProperties = original.RegistrationParameterProperties
type RegistrationProperties = original.RegistrationProperties
type Resource = original.Resource
type URI = original.URI
type VirtualMachineExtensionProductProperties = original.VirtualMachineExtensionProductProperties
type VirtualMachineProductProperties = original.VirtualMachineProductProperties
type OperationsClient = original.OperationsClient
type ProductsClient = original.ProductsClient
type RegistrationsClient = original.RegistrationsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewCustomerSubscriptionsClient(subscriptionID string) CustomerSubscriptionsClient {
	return original.NewCustomerSubscriptionsClient(subscriptionID)
}
func NewCustomerSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) CustomerSubscriptionsClient {
	return original.NewCustomerSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleComputeRoleValues() []ComputeRole {
	return original.PossibleComputeRoleValues()
}
func PossibleOperatingSystemValues() []OperatingSystem {
	return original.PossibleOperatingSystemValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProductsClient(subscriptionID string) ProductsClient {
	return original.NewProductsClient(subscriptionID)
}
func NewProductsClientWithBaseURI(baseURI string, subscriptionID string) ProductsClient {
	return original.NewProductsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegistrationsClient(subscriptionID string) RegistrationsClient {
	return original.NewRegistrationsClient(subscriptionID)
}
func NewRegistrationsClientWithBaseURI(baseURI string, subscriptionID string) RegistrationsClient {
	return original.NewRegistrationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
