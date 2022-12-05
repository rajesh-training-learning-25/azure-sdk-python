//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azsecrets

// BackupSecretResponse contains the response from method Client.BackupSecret.
type BackupSecretResponse struct {
	BackupSecretResult
}

// DeleteSecretResponse contains the response from method Client.DeleteSecret.
type DeleteSecretResponse struct {
	DeletedSecretBundle
}

// GetDeletedSecretResponse contains the response from method Client.GetDeletedSecret.
type GetDeletedSecretResponse struct {
	DeletedSecretBundle
}

// GetSecretResponse contains the response from method Client.GetSecret.
type GetSecretResponse struct {
	SecretBundle
}

// ListDeletedSecretsResponse contains the response from method Client.ListDeletedSecrets.
type ListDeletedSecretsResponse struct {
	DeletedSecretListResult
}

// ListSecretVersionsResponse contains the response from method Client.ListSecretVersions.
type ListSecretVersionsResponse struct {
	SecretListResult
}

// ListSecretsResponse contains the response from method Client.ListSecrets.
type ListSecretsResponse struct {
	SecretListResult
}

// PurgeDeletedSecretResponse contains the response from method Client.PurgeDeletedSecret.
type PurgeDeletedSecretResponse struct {
	// placeholder for future response values
}

// RecoverDeletedSecretResponse contains the response from method Client.RecoverDeletedSecret.
type RecoverDeletedSecretResponse struct {
	SecretBundle
}

// RestoreSecretResponse contains the response from method Client.RestoreSecret.
type RestoreSecretResponse struct {
	SecretBundle
}

// SetSecretResponse contains the response from method Client.SetSecret.
type SetSecretResponse struct {
	SecretBundle
}

// UpdateSecretResponse contains the response from method Client.UpdateSecret.
type UpdateSecretResponse struct {
	SecretBundle
}