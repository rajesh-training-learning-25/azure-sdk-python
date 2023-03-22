//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package rbac

// DataAction - Supported permissions for data actions.
type DataAction string

const (
	// DataActionBackupHsmKeys - Backup HSM keys.
	DataActionBackupHsmKeys DataAction = "Microsoft.KeyVault/managedHsm/keys/backup/action"
	// DataActionCreateHsmKey - Create an HSM key.
	DataActionCreateHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/create"
	// DataActionDecryptHsmKey - Decrypt using an HSM key.
	DataActionDecryptHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/decrypt/action"
	// DataActionDeleteHsmKey - Delete an HSM key.
	DataActionDeleteHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/delete"
	// DataActionDeleteRoleAssignment - Delete role assignment.
	DataActionDeleteRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/delete/action"
	// DataActionDeleteRoleDefinition - Delete role definition.
	DataActionDeleteRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/delete/action"
	// DataActionDownloadHsmSecurityDomain - Download an HSM security domain.
	DataActionDownloadHsmSecurityDomain DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/action"
	// DataActionDownloadHsmSecurityDomainStatus - Check status of HSM security domain download.
	DataActionDownloadHsmSecurityDomainStatus DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/read"
	// DataActionEncryptHsmKey - Encrypt using an HSM key.
	DataActionEncryptHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/encrypt/action"
	// DataActionExportHsmKey - Export an HSM key.
	DataActionExportHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/export/action"
	// DataActionGetRoleAssignment - Get role assignment.
	DataActionGetRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/read/action"
	// DataActionImportHsmKey - Import an HSM key.
	DataActionImportHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/import/action"
	// DataActionPurgeDeletedHsmKey - Purge a deleted HSM key.
	DataActionPurgeDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/delete"
	// DataActionRandomNumbersGenerate - Generate random numbers.
	DataActionRandomNumbersGenerate DataAction = "Microsoft.KeyVault/managedHsm/rng/action"
	// DataActionReadDeletedHsmKey - Read deleted HSM key.
	DataActionReadDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/read/action"
	// DataActionReadHsmBackupStatus - Read an HSM backup status.
	DataActionReadHsmBackupStatus DataAction = "Microsoft.KeyVault/managedHsm/backup/status/action"
	// DataActionReadHsmKey - Read HSM key metadata.
	DataActionReadHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/read/action"
	// DataActionReadHsmRestoreStatus - Read an HSM restore status.
	DataActionReadHsmRestoreStatus DataAction = "Microsoft.KeyVault/managedHsm/restore/status/action"
	// DataActionReadHsmSecurityDomainStatus - Check the status of the HSM security domain exchange file.
	DataActionReadHsmSecurityDomainStatus DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/read"
	// DataActionReadHsmSecurityDomainTransferKey - Download an HSM security domain transfer key.
	DataActionReadHsmSecurityDomainTransferKey DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/transferkey/read"
	// DataActionReadRoleDefinition - Get role definition.
	DataActionReadRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/read/action"
	// DataActionRecoverDeletedHsmKey - Recover deleted HSM key.
	DataActionRecoverDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/recover/action"
	// DataActionReleaseKey - Release an HSM key using Secure Key Release.
	DataActionReleaseKey DataAction = "Microsoft.KeyVault/managedHsm/keys/release/action"
	// DataActionRestoreHsmKeys - Restore HSM keys.
	DataActionRestoreHsmKeys DataAction = "Microsoft.KeyVault/managedHsm/keys/restore/action"
	// DataActionSignHsmKey - Sign using an HSM key.
	DataActionSignHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/sign/action"
	// DataActionStartHsmBackup - Start an HSM backup.
	DataActionStartHsmBackup DataAction = "Microsoft.KeyVault/managedHsm/backup/start/action"
	// DataActionStartHsmRestore - Start an HSM restore.
	DataActionStartHsmRestore DataAction = "Microsoft.KeyVault/managedHsm/restore/start/action"
	// DataActionUnwrapHsmKey - Unwrap using an HSM key.
	DataActionUnwrapHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/unwrap/action"
	// DataActionUploadHsmSecurityDomain - Upload an HSM security domain.
	DataActionUploadHsmSecurityDomain DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/action"
	// DataActionVerifyHsmKey - Verify using an HSM key.
	DataActionVerifyHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/verify/action"
	// DataActionWrapHsmKey - Wrap using an HSM key.
	DataActionWrapHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/wrap/action"
	// DataActionWriteHsmKey - Update an HSM key.
	DataActionWriteHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/write/action"
	// DataActionWriteRoleAssignment - Create or update role assignment.
	DataActionWriteRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/write/action"
	// DataActionWriteRoleDefinition - Create or update role definition.
	DataActionWriteRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/write/action"
)

// PossibleDataActionValues returns the possible values for the DataAction const type.
func PossibleDataActionValues() []DataAction {
	return []DataAction{
		DataActionBackupHsmKeys,
		DataActionCreateHsmKey,
		DataActionDecryptHsmKey,
		DataActionDeleteHsmKey,
		DataActionDeleteRoleAssignment,
		DataActionDeleteRoleDefinition,
		DataActionDownloadHsmSecurityDomain,
		DataActionDownloadHsmSecurityDomainStatus,
		DataActionEncryptHsmKey,
		DataActionExportHsmKey,
		DataActionGetRoleAssignment,
		DataActionImportHsmKey,
		DataActionPurgeDeletedHsmKey,
		DataActionRandomNumbersGenerate,
		DataActionReadDeletedHsmKey,
		DataActionReadHsmBackupStatus,
		DataActionReadHsmKey,
		DataActionReadHsmRestoreStatus,
		DataActionReadHsmSecurityDomainStatus,
		DataActionReadHsmSecurityDomainTransferKey,
		DataActionReadRoleDefinition,
		DataActionRecoverDeletedHsmKey,
		DataActionReleaseKey,
		DataActionRestoreHsmKeys,
		DataActionSignHsmKey,
		DataActionStartHsmBackup,
		DataActionStartHsmRestore,
		DataActionUnwrapHsmKey,
		DataActionUploadHsmSecurityDomain,
		DataActionVerifyHsmKey,
		DataActionWrapHsmKey,
		DataActionWriteHsmKey,
		DataActionWriteRoleAssignment,
		DataActionWriteRoleDefinition,
	}
}

// RoleDefinitionType - The role definition type.
type RoleDefinitionType string

const (
	RoleDefinitionTypeMicrosoftAuthorizationRoleDefinitions RoleDefinitionType = "Microsoft.Authorization/roleDefinitions"
)

// PossibleRoleDefinitionTypeValues returns the possible values for the RoleDefinitionType const type.
func PossibleRoleDefinitionTypeValues() []RoleDefinitionType {
	return []RoleDefinitionType{
		RoleDefinitionTypeMicrosoftAuthorizationRoleDefinitions,
	}
}

// RoleScope - The role scope.
type RoleScope string

const (
	// RoleScopeGlobal - Global scope
	RoleScopeGlobal RoleScope = "/"
	// RoleScopeKeys - Keys scope
	RoleScopeKeys RoleScope = "/keys"
)

// PossibleRoleScopeValues returns the possible values for the RoleScope const type.
func PossibleRoleScopeValues() []RoleScope {
	return []RoleScope{
		RoleScopeGlobal,
		RoleScopeKeys,
	}
}

// RoleType - The role type.
type RoleType string

const (
	// RoleTypeBuiltInRole - Built in role.
	RoleTypeBuiltInRole RoleType = "AKVBuiltInRole"
	// RoleTypeCustomRole - Custom role.
	RoleTypeCustomRole RoleType = "CustomRole"
)

// PossibleRoleTypeValues returns the possible values for the RoleType const type.
func PossibleRoleTypeValues() []RoleType {
	return []RoleType{
		RoleTypeBuiltInRole,
		RoleTypeCustomRole,
	}
}
