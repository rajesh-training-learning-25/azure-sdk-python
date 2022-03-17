//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package keyvault

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/preview/keyvault/v7.2-preview/keyvault"
)

type ActionType = original.ActionType

const (
	AutoRenew     ActionType = original.AutoRenew
	EmailContacts ActionType = original.EmailContacts
)

type DeletionRecoveryLevel = original.DeletionRecoveryLevel

const (
	CustomizedRecoverable                      DeletionRecoveryLevel = original.CustomizedRecoverable
	CustomizedRecoverableProtectedSubscription DeletionRecoveryLevel = original.CustomizedRecoverableProtectedSubscription
	CustomizedRecoverablePurgeable             DeletionRecoveryLevel = original.CustomizedRecoverablePurgeable
	Purgeable                                  DeletionRecoveryLevel = original.Purgeable
	Recoverable                                DeletionRecoveryLevel = original.Recoverable
	RecoverableProtectedSubscription           DeletionRecoveryLevel = original.RecoverableProtectedSubscription
	RecoverablePurgeable                       DeletionRecoveryLevel = original.RecoverablePurgeable
)

type JSONWebKeyCurveName = original.JSONWebKeyCurveName

const (
	P256  JSONWebKeyCurveName = original.P256
	P256K JSONWebKeyCurveName = original.P256K
	P384  JSONWebKeyCurveName = original.P384
	P521  JSONWebKeyCurveName = original.P521
)

type JSONWebKeyEncryptionAlgorithm = original.JSONWebKeyEncryptionAlgorithm

const (
	A128CBC    JSONWebKeyEncryptionAlgorithm = original.A128CBC
	A128CBCPAD JSONWebKeyEncryptionAlgorithm = original.A128CBCPAD
	A128GCM    JSONWebKeyEncryptionAlgorithm = original.A128GCM
	A128KW     JSONWebKeyEncryptionAlgorithm = original.A128KW
	A192CBC    JSONWebKeyEncryptionAlgorithm = original.A192CBC
	A192CBCPAD JSONWebKeyEncryptionAlgorithm = original.A192CBCPAD
	A192GCM    JSONWebKeyEncryptionAlgorithm = original.A192GCM
	A192KW     JSONWebKeyEncryptionAlgorithm = original.A192KW
	A256CBC    JSONWebKeyEncryptionAlgorithm = original.A256CBC
	A256CBCPAD JSONWebKeyEncryptionAlgorithm = original.A256CBCPAD
	A256GCM    JSONWebKeyEncryptionAlgorithm = original.A256GCM
	A256KW     JSONWebKeyEncryptionAlgorithm = original.A256KW
	RSA15      JSONWebKeyEncryptionAlgorithm = original.RSA15
	RSAOAEP    JSONWebKeyEncryptionAlgorithm = original.RSAOAEP
	RSAOAEP256 JSONWebKeyEncryptionAlgorithm = original.RSAOAEP256
)

type JSONWebKeyOperation = original.JSONWebKeyOperation

const (
	Decrypt   JSONWebKeyOperation = original.Decrypt
	Encrypt   JSONWebKeyOperation = original.Encrypt
	Export    JSONWebKeyOperation = original.Export
	Import    JSONWebKeyOperation = original.Import
	Sign      JSONWebKeyOperation = original.Sign
	UnwrapKey JSONWebKeyOperation = original.UnwrapKey
	Verify    JSONWebKeyOperation = original.Verify
	WrapKey   JSONWebKeyOperation = original.WrapKey
)

type JSONWebKeySignatureAlgorithm = original.JSONWebKeySignatureAlgorithm

const (
	ES256  JSONWebKeySignatureAlgorithm = original.ES256
	ES256K JSONWebKeySignatureAlgorithm = original.ES256K
	ES384  JSONWebKeySignatureAlgorithm = original.ES384
	ES512  JSONWebKeySignatureAlgorithm = original.ES512
	PS256  JSONWebKeySignatureAlgorithm = original.PS256
	PS384  JSONWebKeySignatureAlgorithm = original.PS384
	PS512  JSONWebKeySignatureAlgorithm = original.PS512
	RS256  JSONWebKeySignatureAlgorithm = original.RS256
	RS384  JSONWebKeySignatureAlgorithm = original.RS384
	RS512  JSONWebKeySignatureAlgorithm = original.RS512
	RSNULL JSONWebKeySignatureAlgorithm = original.RSNULL
)

type JSONWebKeyType = original.JSONWebKeyType

const (
	EC     JSONWebKeyType = original.EC
	ECHSM  JSONWebKeyType = original.ECHSM
	Oct    JSONWebKeyType = original.Oct
	OctHSM JSONWebKeyType = original.OctHSM
	RSA    JSONWebKeyType = original.RSA
	RSAHSM JSONWebKeyType = original.RSAHSM
)

type KeyUsageType = original.KeyUsageType

const (
	CRLSign          KeyUsageType = original.CRLSign
	DataEncipherment KeyUsageType = original.DataEncipherment
	DecipherOnly     KeyUsageType = original.DecipherOnly
	DigitalSignature KeyUsageType = original.DigitalSignature
	EncipherOnly     KeyUsageType = original.EncipherOnly
	KeyAgreement     KeyUsageType = original.KeyAgreement
	KeyCertSign      KeyUsageType = original.KeyCertSign
	KeyEncipherment  KeyUsageType = original.KeyEncipherment
	NonRepudiation   KeyUsageType = original.NonRepudiation
)

type OperationStatus = original.OperationStatus

const (
	Failed     OperationStatus = original.Failed
	InProgress OperationStatus = original.InProgress
	Success    OperationStatus = original.Success
)

type SasTokenType = original.SasTokenType

const (
	Account SasTokenType = original.Account
	Service SasTokenType = original.Service
)

type Action = original.Action
type AdministratorDetails = original.AdministratorDetails
type Attributes = original.Attributes
type BackupCertificateResult = original.BackupCertificateResult
type BackupKeyResult = original.BackupKeyResult
type BackupSecretResult = original.BackupSecretResult
type BackupStorageResult = original.BackupStorageResult
type BaseClient = original.BaseClient
type CertificateAttributes = original.CertificateAttributes
type CertificateBundle = original.CertificateBundle
type CertificateCreateParameters = original.CertificateCreateParameters
type CertificateImportParameters = original.CertificateImportParameters
type CertificateInfoObject = original.CertificateInfoObject
type CertificateIssuerItem = original.CertificateIssuerItem
type CertificateIssuerListResult = original.CertificateIssuerListResult
type CertificateIssuerListResultIterator = original.CertificateIssuerListResultIterator
type CertificateIssuerListResultPage = original.CertificateIssuerListResultPage
type CertificateIssuerSetParameters = original.CertificateIssuerSetParameters
type CertificateIssuerUpdateParameters = original.CertificateIssuerUpdateParameters
type CertificateItem = original.CertificateItem
type CertificateListResult = original.CertificateListResult
type CertificateListResultIterator = original.CertificateListResultIterator
type CertificateListResultPage = original.CertificateListResultPage
type CertificateMergeParameters = original.CertificateMergeParameters
type CertificateOperation = original.CertificateOperation
type CertificateOperationUpdateParameter = original.CertificateOperationUpdateParameter
type CertificatePolicy = original.CertificatePolicy
type CertificateRestoreParameters = original.CertificateRestoreParameters
type CertificateUpdateParameters = original.CertificateUpdateParameters
type Contact = original.Contact
type Contacts = original.Contacts
type DeletedCertificateBundle = original.DeletedCertificateBundle
type DeletedCertificateItem = original.DeletedCertificateItem
type DeletedCertificateListResult = original.DeletedCertificateListResult
type DeletedCertificateListResultIterator = original.DeletedCertificateListResultIterator
type DeletedCertificateListResultPage = original.DeletedCertificateListResultPage
type DeletedKeyBundle = original.DeletedKeyBundle
type DeletedKeyItem = original.DeletedKeyItem
type DeletedKeyListResult = original.DeletedKeyListResult
type DeletedKeyListResultIterator = original.DeletedKeyListResultIterator
type DeletedKeyListResultPage = original.DeletedKeyListResultPage
type DeletedSasDefinitionBundle = original.DeletedSasDefinitionBundle
type DeletedSasDefinitionItem = original.DeletedSasDefinitionItem
type DeletedSasDefinitionListResult = original.DeletedSasDefinitionListResult
type DeletedSasDefinitionListResultIterator = original.DeletedSasDefinitionListResultIterator
type DeletedSasDefinitionListResultPage = original.DeletedSasDefinitionListResultPage
type DeletedSecretBundle = original.DeletedSecretBundle
type DeletedSecretItem = original.DeletedSecretItem
type DeletedSecretListResult = original.DeletedSecretListResult
type DeletedSecretListResultIterator = original.DeletedSecretListResultIterator
type DeletedSecretListResultPage = original.DeletedSecretListResultPage
type DeletedStorageAccountItem = original.DeletedStorageAccountItem
type DeletedStorageBundle = original.DeletedStorageBundle
type DeletedStorageListResult = original.DeletedStorageListResult
type DeletedStorageListResultIterator = original.DeletedStorageListResultIterator
type DeletedStorageListResultPage = original.DeletedStorageListResultPage
type EncDataSet = original.EncDataSet
type EncDataSetItem = original.EncDataSetItem
type Error = original.Error
type ErrorType = original.ErrorType
type FullBackupFuture = original.FullBackupFuture
type FullBackupOperation = original.FullBackupOperation
type FullRestoreOperationFuture = original.FullRestoreOperationFuture
type HSMSecurityDomainClient = original.HSMSecurityDomainClient
type HSMSecurityDomainUploadFuture = original.HSMSecurityDomainUploadFuture
type IssuerAttributes = original.IssuerAttributes
type IssuerBundle = original.IssuerBundle
type IssuerCredentials = original.IssuerCredentials
type IssuerParameters = original.IssuerParameters
type JSONWebKey = original.JSONWebKey
type Key = original.Key
type KeyAttributes = original.KeyAttributes
type KeyBundle = original.KeyBundle
type KeyCreateParameters = original.KeyCreateParameters
type KeyExportParameters = original.KeyExportParameters
type KeyImportParameters = original.KeyImportParameters
type KeyItem = original.KeyItem
type KeyListResult = original.KeyListResult
type KeyListResultIterator = original.KeyListResultIterator
type KeyListResultPage = original.KeyListResultPage
type KeyOperationResult = original.KeyOperationResult
type KeyOperationsParameters = original.KeyOperationsParameters
type KeyProperties = original.KeyProperties
type KeyReleasePolicy = original.KeyReleasePolicy
type KeyRestoreParameters = original.KeyRestoreParameters
type KeySignParameters = original.KeySignParameters
type KeyUpdateParameters = original.KeyUpdateParameters
type KeyVerifyParameters = original.KeyVerifyParameters
type KeyVerifyResult = original.KeyVerifyResult
type LifetimeAction = original.LifetimeAction
type OrganizationDetails = original.OrganizationDetails
type PendingCertificateSigningRequestResult = original.PendingCertificateSigningRequestResult
type Permission = original.Permission
type RestoreOperation = original.RestoreOperation
type RestoreOperationParameters = original.RestoreOperationParameters
type RoleAssignment = original.RoleAssignment
type RoleAssignmentCreateParameters = original.RoleAssignmentCreateParameters
type RoleAssignmentFilter = original.RoleAssignmentFilter
type RoleAssignmentListResult = original.RoleAssignmentListResult
type RoleAssignmentListResultIterator = original.RoleAssignmentListResultIterator
type RoleAssignmentListResultPage = original.RoleAssignmentListResultPage
type RoleAssignmentProperties = original.RoleAssignmentProperties
type RoleAssignmentPropertiesWithScope = original.RoleAssignmentPropertiesWithScope
type RoleAssignmentsClient = original.RoleAssignmentsClient
type RoleDefinition = original.RoleDefinition
type RoleDefinitionFilter = original.RoleDefinitionFilter
type RoleDefinitionListResult = original.RoleDefinitionListResult
type RoleDefinitionListResultIterator = original.RoleDefinitionListResultIterator
type RoleDefinitionListResultPage = original.RoleDefinitionListResultPage
type RoleDefinitionProperties = original.RoleDefinitionProperties
type RoleDefinitionsClient = original.RoleDefinitionsClient
type SASTokenParameter = original.SASTokenParameter
type SasDefinitionAttributes = original.SasDefinitionAttributes
type SasDefinitionBundle = original.SasDefinitionBundle
type SasDefinitionCreateParameters = original.SasDefinitionCreateParameters
type SasDefinitionItem = original.SasDefinitionItem
type SasDefinitionListResult = original.SasDefinitionListResult
type SasDefinitionListResultIterator = original.SasDefinitionListResultIterator
type SasDefinitionListResultPage = original.SasDefinitionListResultPage
type SasDefinitionUpdateParameters = original.SasDefinitionUpdateParameters
type SecretAttributes = original.SecretAttributes
type SecretBundle = original.SecretBundle
type SecretItem = original.SecretItem
type SecretListResult = original.SecretListResult
type SecretListResultIterator = original.SecretListResultIterator
type SecretListResultPage = original.SecretListResultPage
type SecretProperties = original.SecretProperties
type SecretRestoreParameters = original.SecretRestoreParameters
type SecretSetParameters = original.SecretSetParameters
type SecretUpdateParameters = original.SecretUpdateParameters
type SecurityDomainCertificateItem = original.SecurityDomainCertificateItem
type SecurityDomainJSONWebKey = original.SecurityDomainJSONWebKey
type SecurityDomainObject = original.SecurityDomainObject
type SecurityDomainObjectData = original.SecurityDomainObjectData
type SecurityDomainObjectDataSharedKeys = original.SecurityDomainObjectDataSharedKeys
type SecurityDomainOperationStatus = original.SecurityDomainOperationStatus
type SecurityDomainUploadObject = original.SecurityDomainUploadObject
type SecurityDomainUploadObjectValue = original.SecurityDomainUploadObjectValue
type SecurityDomainUploadObjectValueWrappedKey = original.SecurityDomainUploadObjectValueWrappedKey
type SelectiveKeyRestoreOperation = original.SelectiveKeyRestoreOperation
type SelectiveKeyRestoreOperationMethodFuture = original.SelectiveKeyRestoreOperationMethodFuture
type SelectiveKeyRestoreOperationParameters = original.SelectiveKeyRestoreOperationParameters
type StorageAccountAttributes = original.StorageAccountAttributes
type StorageAccountCreateParameters = original.StorageAccountCreateParameters
type StorageAccountItem = original.StorageAccountItem
type StorageAccountRegenerteKeyParameters = original.StorageAccountRegenerteKeyParameters
type StorageAccountUpdateParameters = original.StorageAccountUpdateParameters
type StorageBundle = original.StorageBundle
type StorageListResult = original.StorageListResult
type StorageListResultIterator = original.StorageListResultIterator
type StorageListResultPage = original.StorageListResultPage
type StorageRestoreParameters = original.StorageRestoreParameters
type SubjectAlternativeNames = original.SubjectAlternativeNames
type TransferKey = original.TransferKey
type Trigger = original.Trigger
type X509CertificateProperties = original.X509CertificateProperties

func New() BaseClient {
	return original.New()
}
func NewCertificateIssuerListResultIterator(page CertificateIssuerListResultPage) CertificateIssuerListResultIterator {
	return original.NewCertificateIssuerListResultIterator(page)
}
func NewCertificateIssuerListResultPage(cur CertificateIssuerListResult, getNextPage func(context.Context, CertificateIssuerListResult) (CertificateIssuerListResult, error)) CertificateIssuerListResultPage {
	return original.NewCertificateIssuerListResultPage(cur, getNextPage)
}
func NewCertificateListResultIterator(page CertificateListResultPage) CertificateListResultIterator {
	return original.NewCertificateListResultIterator(page)
}
func NewCertificateListResultPage(cur CertificateListResult, getNextPage func(context.Context, CertificateListResult) (CertificateListResult, error)) CertificateListResultPage {
	return original.NewCertificateListResultPage(cur, getNextPage)
}
func NewDeletedCertificateListResultIterator(page DeletedCertificateListResultPage) DeletedCertificateListResultIterator {
	return original.NewDeletedCertificateListResultIterator(page)
}
func NewDeletedCertificateListResultPage(cur DeletedCertificateListResult, getNextPage func(context.Context, DeletedCertificateListResult) (DeletedCertificateListResult, error)) DeletedCertificateListResultPage {
	return original.NewDeletedCertificateListResultPage(cur, getNextPage)
}
func NewDeletedKeyListResultIterator(page DeletedKeyListResultPage) DeletedKeyListResultIterator {
	return original.NewDeletedKeyListResultIterator(page)
}
func NewDeletedKeyListResultPage(cur DeletedKeyListResult, getNextPage func(context.Context, DeletedKeyListResult) (DeletedKeyListResult, error)) DeletedKeyListResultPage {
	return original.NewDeletedKeyListResultPage(cur, getNextPage)
}
func NewDeletedSasDefinitionListResultIterator(page DeletedSasDefinitionListResultPage) DeletedSasDefinitionListResultIterator {
	return original.NewDeletedSasDefinitionListResultIterator(page)
}
func NewDeletedSasDefinitionListResultPage(cur DeletedSasDefinitionListResult, getNextPage func(context.Context, DeletedSasDefinitionListResult) (DeletedSasDefinitionListResult, error)) DeletedSasDefinitionListResultPage {
	return original.NewDeletedSasDefinitionListResultPage(cur, getNextPage)
}
func NewDeletedSecretListResultIterator(page DeletedSecretListResultPage) DeletedSecretListResultIterator {
	return original.NewDeletedSecretListResultIterator(page)
}
func NewDeletedSecretListResultPage(cur DeletedSecretListResult, getNextPage func(context.Context, DeletedSecretListResult) (DeletedSecretListResult, error)) DeletedSecretListResultPage {
	return original.NewDeletedSecretListResultPage(cur, getNextPage)
}
func NewDeletedStorageListResultIterator(page DeletedStorageListResultPage) DeletedStorageListResultIterator {
	return original.NewDeletedStorageListResultIterator(page)
}
func NewDeletedStorageListResultPage(cur DeletedStorageListResult, getNextPage func(context.Context, DeletedStorageListResult) (DeletedStorageListResult, error)) DeletedStorageListResultPage {
	return original.NewDeletedStorageListResultPage(cur, getNextPage)
}
func NewHSMSecurityDomainClient() HSMSecurityDomainClient {
	return original.NewHSMSecurityDomainClient()
}
func NewKeyListResultIterator(page KeyListResultPage) KeyListResultIterator {
	return original.NewKeyListResultIterator(page)
}
func NewKeyListResultPage(cur KeyListResult, getNextPage func(context.Context, KeyListResult) (KeyListResult, error)) KeyListResultPage {
	return original.NewKeyListResultPage(cur, getNextPage)
}
func NewRoleAssignmentListResultIterator(page RoleAssignmentListResultPage) RoleAssignmentListResultIterator {
	return original.NewRoleAssignmentListResultIterator(page)
}
func NewRoleAssignmentListResultPage(cur RoleAssignmentListResult, getNextPage func(context.Context, RoleAssignmentListResult) (RoleAssignmentListResult, error)) RoleAssignmentListResultPage {
	return original.NewRoleAssignmentListResultPage(cur, getNextPage)
}
func NewRoleAssignmentsClient() RoleAssignmentsClient {
	return original.NewRoleAssignmentsClient()
}
func NewRoleDefinitionListResultIterator(page RoleDefinitionListResultPage) RoleDefinitionListResultIterator {
	return original.NewRoleDefinitionListResultIterator(page)
}
func NewRoleDefinitionListResultPage(cur RoleDefinitionListResult, getNextPage func(context.Context, RoleDefinitionListResult) (RoleDefinitionListResult, error)) RoleDefinitionListResultPage {
	return original.NewRoleDefinitionListResultPage(cur, getNextPage)
}
func NewRoleDefinitionsClient() RoleDefinitionsClient {
	return original.NewRoleDefinitionsClient()
}
func NewSasDefinitionListResultIterator(page SasDefinitionListResultPage) SasDefinitionListResultIterator {
	return original.NewSasDefinitionListResultIterator(page)
}
func NewSasDefinitionListResultPage(cur SasDefinitionListResult, getNextPage func(context.Context, SasDefinitionListResult) (SasDefinitionListResult, error)) SasDefinitionListResultPage {
	return original.NewSasDefinitionListResultPage(cur, getNextPage)
}
func NewSecretListResultIterator(page SecretListResultPage) SecretListResultIterator {
	return original.NewSecretListResultIterator(page)
}
func NewSecretListResultPage(cur SecretListResult, getNextPage func(context.Context, SecretListResult) (SecretListResult, error)) SecretListResultPage {
	return original.NewSecretListResultPage(cur, getNextPage)
}
func NewStorageListResultIterator(page StorageListResultPage) StorageListResultIterator {
	return original.NewStorageListResultIterator(page)
}
func NewStorageListResultPage(cur StorageListResult, getNextPage func(context.Context, StorageListResult) (StorageListResult, error)) StorageListResultPage {
	return original.NewStorageListResultPage(cur, getNextPage)
}
func NewWithoutDefaults() BaseClient {
	return original.NewWithoutDefaults()
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleDeletionRecoveryLevelValues() []DeletionRecoveryLevel {
	return original.PossibleDeletionRecoveryLevelValues()
}
func PossibleJSONWebKeyCurveNameValues() []JSONWebKeyCurveName {
	return original.PossibleJSONWebKeyCurveNameValues()
}
func PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm {
	return original.PossibleJSONWebKeyEncryptionAlgorithmValues()
}
func PossibleJSONWebKeyOperationValues() []JSONWebKeyOperation {
	return original.PossibleJSONWebKeyOperationValues()
}
func PossibleJSONWebKeySignatureAlgorithmValues() []JSONWebKeySignatureAlgorithm {
	return original.PossibleJSONWebKeySignatureAlgorithmValues()
}
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return original.PossibleJSONWebKeyTypeValues()
}
func PossibleKeyUsageTypeValues() []KeyUsageType {
	return original.PossibleKeyUsageTypeValues()
}
func PossibleOperationStatusValues() []OperationStatus {
	return original.PossibleOperationStatusValues()
}
func PossibleSasTokenTypeValues() []SasTokenType {
	return original.PossibleSasTokenTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
