//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

const (
	moduleName    = "armresources"
	moduleVersion = "v1.0.0"
)

// AliasPathAttributes - The attributes of the token that the alias path is referring to.
type AliasPathAttributes string

const (
	// AliasPathAttributesModifiable - The token that the alias path is referring to is modifiable by policies with 'modify' effect.
	AliasPathAttributesModifiable AliasPathAttributes = "Modifiable"
	// AliasPathAttributesNone - The token that the alias path is referring to has no attributes.
	AliasPathAttributesNone AliasPathAttributes = "None"
)

// PossibleAliasPathAttributesValues returns the possible values for the AliasPathAttributes const type.
func PossibleAliasPathAttributesValues() []AliasPathAttributes {
	return []AliasPathAttributes{
		AliasPathAttributesModifiable,
		AliasPathAttributesNone,
	}
}

// AliasPathTokenType - The type of the token that the alias path is referring to.
type AliasPathTokenType string

const (
	// AliasPathTokenTypeAny - The token type can be anything.
	AliasPathTokenTypeAny AliasPathTokenType = "Any"
	// AliasPathTokenTypeArray - The token type is array.
	AliasPathTokenTypeArray AliasPathTokenType = "Array"
	// AliasPathTokenTypeBoolean - The token type is boolean.
	AliasPathTokenTypeBoolean AliasPathTokenType = "Boolean"
	// AliasPathTokenTypeInteger - The token type is integer.
	AliasPathTokenTypeInteger AliasPathTokenType = "Integer"
	// AliasPathTokenTypeNotSpecified - The token type is not specified.
	AliasPathTokenTypeNotSpecified AliasPathTokenType = "NotSpecified"
	// AliasPathTokenTypeNumber - The token type is number.
	AliasPathTokenTypeNumber AliasPathTokenType = "Number"
	// AliasPathTokenTypeObject - The token type is object.
	AliasPathTokenTypeObject AliasPathTokenType = "Object"
	// AliasPathTokenTypeString - The token type is string.
	AliasPathTokenTypeString AliasPathTokenType = "String"
)

// PossibleAliasPathTokenTypeValues returns the possible values for the AliasPathTokenType const type.
func PossibleAliasPathTokenTypeValues() []AliasPathTokenType {
	return []AliasPathTokenType{
		AliasPathTokenTypeAny,
		AliasPathTokenTypeArray,
		AliasPathTokenTypeBoolean,
		AliasPathTokenTypeInteger,
		AliasPathTokenTypeNotSpecified,
		AliasPathTokenTypeNumber,
		AliasPathTokenTypeObject,
		AliasPathTokenTypeString,
	}
}

// AliasPatternType - The type of alias pattern
type AliasPatternType string

const (
	// AliasPatternTypeNotSpecified - NotSpecified is not allowed.
	AliasPatternTypeNotSpecified AliasPatternType = "NotSpecified"
	// AliasPatternTypeExtract - Extract is the only allowed value.
	AliasPatternTypeExtract AliasPatternType = "Extract"
)

// PossibleAliasPatternTypeValues returns the possible values for the AliasPatternType const type.
func PossibleAliasPatternTypeValues() []AliasPatternType {
	return []AliasPatternType{
		AliasPatternTypeNotSpecified,
		AliasPatternTypeExtract,
	}
}

// AliasType - The type of the alias.
type AliasType string

const (
	// AliasTypeNotSpecified - Alias type is unknown (same as not providing alias type).
	AliasTypeNotSpecified AliasType = "NotSpecified"
	// AliasTypePlainText - Alias value is not secret.
	AliasTypePlainText AliasType = "PlainText"
	// AliasTypeMask - Alias value is secret.
	AliasTypeMask AliasType = "Mask"
)

// PossibleAliasTypeValues returns the possible values for the AliasType const type.
func PossibleAliasTypeValues() []AliasType {
	return []AliasType{
		AliasTypeNotSpecified,
		AliasTypePlainText,
		AliasTypeMask,
	}
}

// ChangeType - Type of change that will be made to the resource when the deployment is executed.
type ChangeType string

const (
	// ChangeTypeCreate - The resource does not exist in the current state but is present in the desired state. The resource will
	// be created when the deployment is executed.
	ChangeTypeCreate ChangeType = "Create"
	// ChangeTypeDelete - The resource exists in the current state and is missing from the desired state. The resource will be
	// deleted when the deployment is executed.
	ChangeTypeDelete ChangeType = "Delete"
	// ChangeTypeIgnore - The resource exists in the current state and is missing from the desired state. The resource will not
	// be deployed or modified when the deployment is executed.
	ChangeTypeIgnore ChangeType = "Ignore"
	// ChangeTypeDeploy - The resource exists in the current state and the desired state and will be redeployed when the deployment
	// is executed. The properties of the resource may or may not change.
	ChangeTypeDeploy ChangeType = "Deploy"
	// ChangeTypeNoChange - The resource exists in the current state and the desired state and will be redeployed when the deployment
	// is executed. The properties of the resource will not change.
	ChangeTypeNoChange ChangeType = "NoChange"
	// ChangeTypeModify - The resource exists in the current state and the desired state and will be redeployed when the deployment
	// is executed. The properties of the resource will change.
	ChangeTypeModify ChangeType = "Modify"
	// ChangeTypeUnsupported - The resource is not supported by What-If.
	ChangeTypeUnsupported ChangeType = "Unsupported"
)

// PossibleChangeTypeValues returns the possible values for the ChangeType const type.
func PossibleChangeTypeValues() []ChangeType {
	return []ChangeType{
		ChangeTypeCreate,
		ChangeTypeDelete,
		ChangeTypeIgnore,
		ChangeTypeDeploy,
		ChangeTypeNoChange,
		ChangeTypeModify,
		ChangeTypeUnsupported,
	}
}

// DeploymentMode - The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental
// mode, resources are deployed without deleting existing resources that are not included in
// the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included
// in the template are deleted. Be careful when using Complete mode as you may
// unintentionally delete resources.
type DeploymentMode string

const (
	DeploymentModeIncremental DeploymentMode = "Incremental"
	DeploymentModeComplete    DeploymentMode = "Complete"
)

// PossibleDeploymentModeValues returns the possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{
		DeploymentModeIncremental,
		DeploymentModeComplete,
	}
}

// ExpressionEvaluationOptionsScopeType - The scope to be used for evaluation of parameters, variables and functions in a
// nested template.
type ExpressionEvaluationOptionsScopeType string

const (
	ExpressionEvaluationOptionsScopeTypeInner        ExpressionEvaluationOptionsScopeType = "Inner"
	ExpressionEvaluationOptionsScopeTypeNotSpecified ExpressionEvaluationOptionsScopeType = "NotSpecified"
	ExpressionEvaluationOptionsScopeTypeOuter        ExpressionEvaluationOptionsScopeType = "Outer"
)

// PossibleExpressionEvaluationOptionsScopeTypeValues returns the possible values for the ExpressionEvaluationOptionsScopeType const type.
func PossibleExpressionEvaluationOptionsScopeTypeValues() []ExpressionEvaluationOptionsScopeType {
	return []ExpressionEvaluationOptionsScopeType{
		ExpressionEvaluationOptionsScopeTypeInner,
		ExpressionEvaluationOptionsScopeTypeNotSpecified,
		ExpressionEvaluationOptionsScopeTypeOuter,
	}
}

// ExtendedLocationType - The extended location type.
type ExtendedLocationType string

const (
	ExtendedLocationTypeEdgeZone ExtendedLocationType = "EdgeZone"
)

// PossibleExtendedLocationTypeValues returns the possible values for the ExtendedLocationType const type.
func PossibleExtendedLocationTypeValues() []ExtendedLocationType {
	return []ExtendedLocationType{
		ExtendedLocationTypeEdgeZone,
	}
}

// OnErrorDeploymentType - The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
type OnErrorDeploymentType string

const (
	OnErrorDeploymentTypeLastSuccessful     OnErrorDeploymentType = "LastSuccessful"
	OnErrorDeploymentTypeSpecificDeployment OnErrorDeploymentType = "SpecificDeployment"
)

// PossibleOnErrorDeploymentTypeValues returns the possible values for the OnErrorDeploymentType const type.
func PossibleOnErrorDeploymentTypeValues() []OnErrorDeploymentType {
	return []OnErrorDeploymentType{
		OnErrorDeploymentTypeLastSuccessful,
		OnErrorDeploymentTypeSpecificDeployment,
	}
}

// PropertyChangeType - The type of property change.
type PropertyChangeType string

const (
	// PropertyChangeTypeCreate - The property does not exist in the current state but is present in the desired state. The property
	// will be created when the deployment is executed.
	PropertyChangeTypeCreate PropertyChangeType = "Create"
	// PropertyChangeTypeDelete - The property exists in the current state and is missing from the desired state. It will be deleted
	// when the deployment is executed.
	PropertyChangeTypeDelete PropertyChangeType = "Delete"
	// PropertyChangeTypeModify - The property exists in both current and desired state and is different. The value of the property
	// will change when the deployment is executed.
	PropertyChangeTypeModify PropertyChangeType = "Modify"
	// PropertyChangeTypeArray - The property is an array and contains nested changes.
	PropertyChangeTypeArray PropertyChangeType = "Array"
	// PropertyChangeTypeNoEffect - The property will not be set or updated.
	PropertyChangeTypeNoEffect PropertyChangeType = "NoEffect"
)

// PossiblePropertyChangeTypeValues returns the possible values for the PropertyChangeType const type.
func PossiblePropertyChangeTypeValues() []PropertyChangeType {
	return []PropertyChangeType{
		PropertyChangeTypeCreate,
		PropertyChangeTypeDelete,
		PropertyChangeTypeModify,
		PropertyChangeTypeArray,
		PropertyChangeTypeNoEffect,
	}
}

// ProviderAuthorizationConsentState - The provider authorization consent state.
type ProviderAuthorizationConsentState string

const (
	ProviderAuthorizationConsentStateConsented    ProviderAuthorizationConsentState = "Consented"
	ProviderAuthorizationConsentStateNotRequired  ProviderAuthorizationConsentState = "NotRequired"
	ProviderAuthorizationConsentStateNotSpecified ProviderAuthorizationConsentState = "NotSpecified"
	ProviderAuthorizationConsentStateRequired     ProviderAuthorizationConsentState = "Required"
)

// PossibleProviderAuthorizationConsentStateValues returns the possible values for the ProviderAuthorizationConsentState const type.
func PossibleProviderAuthorizationConsentStateValues() []ProviderAuthorizationConsentState {
	return []ProviderAuthorizationConsentState{
		ProviderAuthorizationConsentStateConsented,
		ProviderAuthorizationConsentStateNotRequired,
		ProviderAuthorizationConsentStateNotSpecified,
		ProviderAuthorizationConsentStateRequired,
	}
}

// ProvisioningOperation - The name of the current provisioning operation.
type ProvisioningOperation string

const (
	// ProvisioningOperationNotSpecified - The provisioning operation is not specified.
	ProvisioningOperationNotSpecified ProvisioningOperation = "NotSpecified"
	// ProvisioningOperationCreate - The provisioning operation is create.
	ProvisioningOperationCreate ProvisioningOperation = "Create"
	// ProvisioningOperationDelete - The provisioning operation is delete.
	ProvisioningOperationDelete ProvisioningOperation = "Delete"
	// ProvisioningOperationWaiting - The provisioning operation is waiting.
	ProvisioningOperationWaiting ProvisioningOperation = "Waiting"
	// ProvisioningOperationAzureAsyncOperationWaiting - The provisioning operation is waiting Azure async operation.
	ProvisioningOperationAzureAsyncOperationWaiting ProvisioningOperation = "AzureAsyncOperationWaiting"
	// ProvisioningOperationResourceCacheWaiting - The provisioning operation is waiting for resource cache.
	ProvisioningOperationResourceCacheWaiting ProvisioningOperation = "ResourceCacheWaiting"
	// ProvisioningOperationAction - The provisioning operation is action.
	ProvisioningOperationAction ProvisioningOperation = "Action"
	// ProvisioningOperationRead - The provisioning operation is read.
	ProvisioningOperationRead ProvisioningOperation = "Read"
	// ProvisioningOperationEvaluateDeploymentOutput - The provisioning operation is evaluate output.
	ProvisioningOperationEvaluateDeploymentOutput ProvisioningOperation = "EvaluateDeploymentOutput"
	// ProvisioningOperationDeploymentCleanup - The provisioning operation is cleanup. This operation is part of the 'complete'
	// mode deployment.
	ProvisioningOperationDeploymentCleanup ProvisioningOperation = "DeploymentCleanup"
)

// PossibleProvisioningOperationValues returns the possible values for the ProvisioningOperation const type.
func PossibleProvisioningOperationValues() []ProvisioningOperation {
	return []ProvisioningOperation{
		ProvisioningOperationNotSpecified,
		ProvisioningOperationCreate,
		ProvisioningOperationDelete,
		ProvisioningOperationWaiting,
		ProvisioningOperationAzureAsyncOperationWaiting,
		ProvisioningOperationResourceCacheWaiting,
		ProvisioningOperationAction,
		ProvisioningOperationRead,
		ProvisioningOperationEvaluateDeploymentOutput,
		ProvisioningOperationDeploymentCleanup,
	}
}

// ProvisioningState - Denotes the state of provisioning.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreated      ProvisioningState = "Created"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateReady        ProvisioningState = "Ready"
	ProvisioningStateRunning      ProvisioningState = "Running"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateReady,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ResourceIdentityType - The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// TagsPatchOperation - The operation type for the patch API.
type TagsPatchOperation string

const (
	// TagsPatchOperationDelete - The 'delete' option allows selectively deleting tags based on given names or name/value pairs.
	TagsPatchOperationDelete TagsPatchOperation = "Delete"
	// TagsPatchOperationMerge - The 'merge' option allows adding tags with new names and updating the values of tags with existing
	// names.
	TagsPatchOperationMerge TagsPatchOperation = "Merge"
	// TagsPatchOperationReplace - The 'replace' option replaces the entire set of existing tags with a new set.
	TagsPatchOperationReplace TagsPatchOperation = "Replace"
)

// PossibleTagsPatchOperationValues returns the possible values for the TagsPatchOperation const type.
func PossibleTagsPatchOperationValues() []TagsPatchOperation {
	return []TagsPatchOperation{
		TagsPatchOperationDelete,
		TagsPatchOperationMerge,
		TagsPatchOperationReplace,
	}
}

// WhatIfResultFormat - The format of the What-If results
type WhatIfResultFormat string

const (
	WhatIfResultFormatResourceIDOnly       WhatIfResultFormat = "ResourceIdOnly"
	WhatIfResultFormatFullResourcePayloads WhatIfResultFormat = "FullResourcePayloads"
)

// PossibleWhatIfResultFormatValues returns the possible values for the WhatIfResultFormat const type.
func PossibleWhatIfResultFormatValues() []WhatIfResultFormat {
	return []WhatIfResultFormat{
		WhatIfResultFormatResourceIDOnly,
		WhatIfResultFormatFullResourcePayloads,
	}
}