//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package generated

// AzureAppConfigurationClientCheckKeyValueResponse contains the response from method AzureAppConfigurationClient.CheckKeyValue.
type AzureAppConfigurationClientCheckKeyValueResponse struct {
	// ETag contains the information returned from the ETag header response.
	ETag *string

	// LastModified contains the information returned from the Last-Modified header response.
	LastModified *string

	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientCheckKeyValuesResponse contains the response from method AzureAppConfigurationClient.CheckKeyValues.
type AzureAppConfigurationClientCheckKeyValuesResponse struct {
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientCheckKeysResponse contains the response from method AzureAppConfigurationClient.CheckKeys.
type AzureAppConfigurationClientCheckKeysResponse struct {
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientCheckLabelsResponse contains the response from method AzureAppConfigurationClient.CheckLabels.
type AzureAppConfigurationClientCheckLabelsResponse struct {
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientCheckRevisionsResponse contains the response from method AzureAppConfigurationClient.CheckRevisions.
type AzureAppConfigurationClientCheckRevisionsResponse struct {
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientDeleteKeyValueResponse contains the response from method AzureAppConfigurationClient.DeleteKeyValue.
type AzureAppConfigurationClientDeleteKeyValueResponse struct {
	KeyValue
	// ETag contains the information returned from the ETag header response.
	ETag *string

	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientDeleteLockResponse contains the response from method AzureAppConfigurationClient.DeleteLock.
type AzureAppConfigurationClientDeleteLockResponse struct {
	KeyValue
	// ETag contains the information returned from the ETag header response.
	ETag *string

	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientGetKeyValueResponse contains the response from method AzureAppConfigurationClient.GetKeyValue.
type AzureAppConfigurationClientGetKeyValueResponse struct {
	KeyValue
	// ETag contains the information returned from the ETag header response.
	ETag *string

	// LastModified contains the information returned from the Last-Modified header response.
	LastModified *string

	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientGetKeyValuesResponse contains the response from method AzureAppConfigurationClient.GetKeyValues.
type AzureAppConfigurationClientGetKeyValuesResponse struct {
	KeyValueListResult
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientGetKeysResponse contains the response from method AzureAppConfigurationClient.GetKeys.
type AzureAppConfigurationClientGetKeysResponse struct {
	KeyListResult
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientGetLabelsResponse contains the response from method AzureAppConfigurationClient.GetLabels.
type AzureAppConfigurationClientGetLabelsResponse struct {
	LabelListResult
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientGetNextPageResponse contains the response from method AzureAppConfigurationClient.GetNextPage.
type AzureAppConfigurationClientGetNextPageResponse struct {
	KeyListResult
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientGetRevisionsResponse contains the response from method AzureAppConfigurationClient.GetRevisions.
type AzureAppConfigurationClientGetRevisionsResponse struct {
	KeyValueListResult
	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientPutKeyValueResponse contains the response from method AzureAppConfigurationClient.PutKeyValue.
type AzureAppConfigurationClientPutKeyValueResponse struct {
	KeyValue
	// ETag contains the information returned from the ETag header response.
	ETag *string

	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}

// AzureAppConfigurationClientPutLockResponse contains the response from method AzureAppConfigurationClient.PutLock.
type AzureAppConfigurationClientPutLockResponse struct {
	KeyValue
	// ETag contains the information returned from the ETag header response.
	ETag *string

	// SyncToken contains the information returned from the Sync-Token header response.
	SyncToken *string
}
