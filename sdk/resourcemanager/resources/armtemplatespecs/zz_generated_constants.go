//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtemplatespecs

const (
	module  = "armtemplatespecs"
	version = "v0.1.0"
)

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

type TemplateSpecExpandKind string

const (
	// TemplateSpecExpandKindVersions - Includes version information with the Template Spec.
	TemplateSpecExpandKindVersions TemplateSpecExpandKind = "versions"
)

// PossibleTemplateSpecExpandKindValues returns the possible values for the TemplateSpecExpandKind const type.
func PossibleTemplateSpecExpandKindValues() []TemplateSpecExpandKind {
	return []TemplateSpecExpandKind{
		TemplateSpecExpandKindVersions,
	}
}

// ToPtr returns a *TemplateSpecExpandKind pointing to the current value.
func (c TemplateSpecExpandKind) ToPtr() *TemplateSpecExpandKind {
	return &c
}
