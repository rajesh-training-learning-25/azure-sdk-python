//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomproviders

const (
	moduleName    = "armcustomproviders"
	moduleVersion = "v0.5.0"
)

// ActionRouting - The routing types that are supported for action requests.
type ActionRouting string

const (
	ActionRoutingProxy ActionRouting = "Proxy"
)

// PossibleActionRoutingValues returns the possible values for the ActionRouting const type.
func PossibleActionRoutingValues() []ActionRouting {
	return []ActionRouting{
		ActionRoutingProxy,
	}
}

// ProvisioningState - The provisioning state of the resource provider.
type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
	}
}

// ResourceTypeRouting - The routing types that are supported for resource requests.
type ResourceTypeRouting string

const (
	ResourceTypeRoutingProxy      ResourceTypeRouting = "Proxy"
	ResourceTypeRoutingProxyCache ResourceTypeRouting = "Proxy,Cache"
)

// PossibleResourceTypeRoutingValues returns the possible values for the ResourceTypeRouting const type.
func PossibleResourceTypeRoutingValues() []ResourceTypeRouting {
	return []ResourceTypeRouting{
		ResourceTypeRoutingProxy,
		ResourceTypeRoutingProxyCache,
	}
}

// ValidationType - The type of validation to run against a matching request.
type ValidationType string

const (
	ValidationTypeSwagger ValidationType = "Swagger"
)

// PossibleValidationTypeValues returns the possible values for the ValidationType const type.
func PossibleValidationTypeValues() []ValidationType {
	return []ValidationType{
		ValidationTypeSwagger,
	}
}