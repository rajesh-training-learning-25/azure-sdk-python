package storagedatalake

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"

// DataLakeStorageError ...
type DataLakeStorageError struct {
	// Error - The service error response object.
	Error *DataLakeStorageErrorError `json:"error,omitempty"`
}

// DataLakeStorageErrorError the service error response object.
type DataLakeStorageErrorError struct {
	// Code - The service error code.
	Code *string `json:"code,omitempty"`
	// Message - The service error message.
	Message *string `json:"message,omitempty"`
}

// Filesystem ...
type Filesystem struct {
	Name         *string `json:"name,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	ETag         *string `json:"eTag,omitempty"`
}

// FilesystemList ...
type FilesystemList struct {
	autorest.Response `json:"-"`
	Filesystems       *[]Filesystem `json:"filesystems,omitempty"`
}

// Path ...
type Path struct {
	Name          *string `json:"name,omitempty"`
	IsDirectory   *string   `json:"isDirectory,omitempty"`
	LastModified  *string `json:"lastModified,omitempty"`
	ETag          *string `json:"eTag,omitempty"`
	ContentLength *string  `json:"contentLength,omitempty"`
	Owner         *string `json:"owner,omitempty"`
	Group         *string `json:"group,omitempty"`
	Permissions   *string `json:"permissions,omitempty"`
}

// PathList ...
type PathList struct {
	autorest.Response `json:"-"`
	Paths             *[]Path `json:"paths,omitempty"`
}

// ReadCloser ...
type ReadCloser struct {
	autorest.Response `json:"-"`
	Value             *io.ReadCloser `json:"value,omitempty"`
}
