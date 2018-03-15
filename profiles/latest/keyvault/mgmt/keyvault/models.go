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

package keyvault

import original "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type AccessPolicyUpdateKind = original.AccessPolicyUpdateKind

const (
	Add     AccessPolicyUpdateKind = original.Add
	Remove  AccessPolicyUpdateKind = original.Remove
	Replace AccessPolicyUpdateKind = original.Replace
)

func PossibleAccessPolicyUpdateKindValues() [3]AccessPolicyUpdateKind {
	return original.PossibleAccessPolicyUpdateKindValues()
}

type CertificatePermissions = original.CertificatePermissions

const (
	Create         CertificatePermissions = original.Create
	Delete         CertificatePermissions = original.Delete
	Deleteissuers  CertificatePermissions = original.Deleteissuers
	Get            CertificatePermissions = original.Get
	Getissuers     CertificatePermissions = original.Getissuers
	Import         CertificatePermissions = original.Import
	List           CertificatePermissions = original.List
	Listissuers    CertificatePermissions = original.Listissuers
	Managecontacts CertificatePermissions = original.Managecontacts
	Manageissuers  CertificatePermissions = original.Manageissuers
	Purge          CertificatePermissions = original.Purge
	Recover        CertificatePermissions = original.Recover
	Setissuers     CertificatePermissions = original.Setissuers
	Update         CertificatePermissions = original.Update
)

func PossibleCertificatePermissionsValues() [14]CertificatePermissions {
	return original.PossibleCertificatePermissionsValues()
}

type CreateMode = original.CreateMode

const (
	CreateModeDefault CreateMode = original.CreateModeDefault
	CreateModeRecover CreateMode = original.CreateModeRecover
)

func PossibleCreateModeValues() [2]CreateMode {
	return original.PossibleCreateModeValues()
}

type KeyPermissions = original.KeyPermissions

const (
	KeyPermissionsBackup    KeyPermissions = original.KeyPermissionsBackup
	KeyPermissionsCreate    KeyPermissions = original.KeyPermissionsCreate
	KeyPermissionsDecrypt   KeyPermissions = original.KeyPermissionsDecrypt
	KeyPermissionsDelete    KeyPermissions = original.KeyPermissionsDelete
	KeyPermissionsEncrypt   KeyPermissions = original.KeyPermissionsEncrypt
	KeyPermissionsGet       KeyPermissions = original.KeyPermissionsGet
	KeyPermissionsImport    KeyPermissions = original.KeyPermissionsImport
	KeyPermissionsList      KeyPermissions = original.KeyPermissionsList
	KeyPermissionsPurge     KeyPermissions = original.KeyPermissionsPurge
	KeyPermissionsRecover   KeyPermissions = original.KeyPermissionsRecover
	KeyPermissionsRestore   KeyPermissions = original.KeyPermissionsRestore
	KeyPermissionsSign      KeyPermissions = original.KeyPermissionsSign
	KeyPermissionsUnwrapKey KeyPermissions = original.KeyPermissionsUnwrapKey
	KeyPermissionsUpdate    KeyPermissions = original.KeyPermissionsUpdate
	KeyPermissionsVerify    KeyPermissions = original.KeyPermissionsVerify
	KeyPermissionsWrapKey   KeyPermissions = original.KeyPermissionsWrapKey
)

func PossibleKeyPermissionsValues() [16]KeyPermissions {
	return original.PossibleKeyPermissionsValues()
}

type Reason = original.Reason

const (
	AccountNameInvalid Reason = original.AccountNameInvalid
	AlreadyExists      Reason = original.AlreadyExists
)

func PossibleReasonValues() [2]Reason {
	return original.PossibleReasonValues()
}

type SecretPermissions = original.SecretPermissions

const (
	SecretPermissionsBackup  SecretPermissions = original.SecretPermissionsBackup
	SecretPermissionsDelete  SecretPermissions = original.SecretPermissionsDelete
	SecretPermissionsGet     SecretPermissions = original.SecretPermissionsGet
	SecretPermissionsList    SecretPermissions = original.SecretPermissionsList
	SecretPermissionsPurge   SecretPermissions = original.SecretPermissionsPurge
	SecretPermissionsRecover SecretPermissions = original.SecretPermissionsRecover
	SecretPermissionsRestore SecretPermissions = original.SecretPermissionsRestore
	SecretPermissionsSet     SecretPermissions = original.SecretPermissionsSet
)

func PossibleSecretPermissionsValues() [8]SecretPermissions {
	return original.PossibleSecretPermissionsValues()
}

type SkuName = original.SkuName

const (
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

func PossibleSkuNameValues() [2]SkuName {
	return original.PossibleSkuNameValues()
}

type StoragePermissions = original.StoragePermissions

const (
	StoragePermissionsBackup        StoragePermissions = original.StoragePermissionsBackup
	StoragePermissionsDelete        StoragePermissions = original.StoragePermissionsDelete
	StoragePermissionsDeletesas     StoragePermissions = original.StoragePermissionsDeletesas
	StoragePermissionsGet           StoragePermissions = original.StoragePermissionsGet
	StoragePermissionsGetsas        StoragePermissions = original.StoragePermissionsGetsas
	StoragePermissionsList          StoragePermissions = original.StoragePermissionsList
	StoragePermissionsListsas       StoragePermissions = original.StoragePermissionsListsas
	StoragePermissionsPurge         StoragePermissions = original.StoragePermissionsPurge
	StoragePermissionsRecover       StoragePermissions = original.StoragePermissionsRecover
	StoragePermissionsRegeneratekey StoragePermissions = original.StoragePermissionsRegeneratekey
	StoragePermissionsRestore       StoragePermissions = original.StoragePermissionsRestore
	StoragePermissionsSet           StoragePermissions = original.StoragePermissionsSet
	StoragePermissionsSetsas        StoragePermissions = original.StoragePermissionsSetsas
	StoragePermissionsUpdate        StoragePermissions = original.StoragePermissionsUpdate
)

func PossibleStoragePermissionsValues() [14]StoragePermissions {
	return original.PossibleStoragePermissionsValues()
}

type AccessPolicyEntry = original.AccessPolicyEntry
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type DeletedVault = original.DeletedVault
type DeletedVaultListResult = original.DeletedVaultListResult
type DeletedVaultListResultIterator = original.DeletedVaultListResultIterator
type DeletedVaultListResultPage = original.DeletedVaultListResultPage
type DeletedVaultProperties = original.DeletedVaultProperties
type LogSpecification = original.LogSpecification
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationProperties = original.OperationProperties
type Permissions = original.Permissions
type Resource = original.Resource
type ResourceListResult = original.ResourceListResult
type ResourceListResultIterator = original.ResourceListResultIterator
type ResourceListResultPage = original.ResourceListResultPage
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type Vault = original.Vault
type VaultAccessPolicyParameters = original.VaultAccessPolicyParameters
type VaultAccessPolicyProperties = original.VaultAccessPolicyProperties
type VaultCheckNameAvailabilityParameters = original.VaultCheckNameAvailabilityParameters
type VaultCreateOrUpdateParameters = original.VaultCreateOrUpdateParameters
type VaultListResult = original.VaultListResult
type VaultListResultIterator = original.VaultListResultIterator
type VaultListResultPage = original.VaultListResultPage
type VaultPatchParameters = original.VaultPatchParameters
type VaultPatchProperties = original.VaultPatchProperties
type VaultProperties = original.VaultProperties
type VaultsPurgeDeletedFuture = original.VaultsPurgeDeletedFuture
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type VaultsClient = original.VaultsClient

func NewVaultsClient(subscriptionID string) VaultsClient {
	return original.NewVaultsClient(subscriptionID)
}
func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
	return original.NewVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
