//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package blockblob

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/lease"
)

type UploadOptions struct {
	// Optional. Used to set blob tags in various blob operations.
	Tags map[string]string

	// Optional. Specifies a user-defined name-value pair associated with the blob.
	Metadata map[string]string

	// Optional. Indicates the tier to be set on the blob.
	Tier *blob.AccessTier

	// Specify the transactional md5 for the body, to be validated by the service.
	TransactionalContentMD5 []byte

	HTTPHeaders      *blob.HTTPHeaders
	CpkInfo          *blob.CpkInfo
	CpkScopeInfo     *blob.CpkScopeInfo
	AccessConditions *blob.AccessConditions
}

type StageBlockOptions struct {
	CpkInfo *blob.CpkInfo

	CpkScopeInfo *blob.CpkScopeInfo

	LeaseAccessConditions *lease.AccessConditions

	// Specify the transactional crc64 for the body, to be validated by the service.
	TransactionalContentCRC64 []byte

	// Specify the transactional md5 for the body, to be validated by the service.
	TransactionalContentMD5 []byte
}

type SourceModifiedAccessConditions = blob.SourceModifiedAccessConditions

// StageBlockFromURLOptions provides set of configurations for StageBlockFromURL operation
type StageBlockFromURLOptions struct {
	// Only Bearer type is supported. Credentials should be a valid OAuth access token to copy source.
	CopySourceAuthorization *string

	LeaseAccessConditions *lease.AccessConditions

	SourceModifiedAccessConditions *SourceModifiedAccessConditions
	// Specify the md5 calculated for the range of bytes that must be read from the copy source.
	SourceContentMD5 []byte
	// Specify the crc64 calculated for the range of bytes that must be read from the copy source.
	SourceContentCRC64 []byte

	Offset *int64

	Count *int64

	CpkInfo *blob.CpkInfo

	CpkScopeInfo *blob.CpkScopeInfo
}

func (o *StageBlockFromURLOptions) format() (*generated.BlockBlobClientStageBlockFromURLOptions, *generated.CpkInfo, *generated.CpkScopeInfo, *generated.LeaseAccessConditions, *generated.SourceModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil, nil
	}

	options := &generated.BlockBlobClientStageBlockFromURLOptions{
		CopySourceAuthorization: o.CopySourceAuthorization,
		SourceContentMD5:        o.SourceContentMD5,
		SourceContentcrc64:      o.SourceContentCRC64,
		SourceRange:             shared.GetSourceRange(o.Offset, o.Count),
	}

	return options, o.CpkInfo, o.CpkScopeInfo, o.LeaseAccessConditions, o.SourceModifiedAccessConditions
}

type CommitBlockListOptions struct {
	Tags                      map[string]string
	Metadata                  map[string]string
	RequestID                 *string
	Tier                      *blob.AccessTier
	Timeout                   *int32
	TransactionalContentCRC64 []byte
	TransactionalContentMD5   []byte
	HTTPHeaders               *blob.HTTPHeaders
	CpkInfo                   *blob.CpkInfo
	CpkScopeInfo              *blob.CpkScopeInfo
	AccessConditions          *blob.AccessConditions
}

// GetBlockListOptions provides set of configurations for GetBlockList operation
type GetBlockListOptions struct {
	Snapshot         *string
	AccessConditions *blob.AccessConditions
}

func (o *GetBlockListOptions) format() (*generated.BlockBlobClientGetBlockListOptions, *generated.LeaseAccessConditions, *generated.ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}

	leaseAccessConditions, modifiedAccessConditions := exported.FormatBlobAccessConditions(o.AccessConditions)
	return &generated.BlockBlobClientGetBlockListOptions{Snapshot: o.Snapshot}, leaseAccessConditions, modifiedAccessConditions
}

// ------------------------------------------------------------

// UploadReaderAtToBlockBlobOption identifies options used by the UploadBuffer and UploadFile functions.
type UploadReaderAtToBlockBlobOption struct {
	// BlockSize specifies the block size to use; the default (and maximum size) is MaxStageBlockBytes.
	BlockSize int64

	// Progress is a function that is invoked periodically as bytes are sent to the BlockBlobClient.
	// Note that the progress reporting is not always increasing; it can go down when retrying a request.
	Progress func(bytesTransferred int64)

	// HTTPHeaders indicates the HTTP headers to be associated with the blob.
	HTTPHeaders *blob.HTTPHeaders

	// Metadata indicates the metadata to be associated with the blob when PutBlockList is called.
	Metadata map[string]string

	// AccessConditions indicates the access conditions for the block blob.
	AccessConditions *blob.AccessConditions

	// AccessTier indicates the tier of blob
	AccessTier *blob.AccessTier

	// BlobTags
	Tags map[string]string

	// ClientProvidedKeyOptions indicates the client provided key by name and/or by value to encrypt/decrypt data.
	CpkInfo      *blob.CpkInfo
	CpkScopeInfo *blob.CpkScopeInfo

	// Parallelism indicates the maximum number of blocks to upload in parallel (0=default)
	Parallelism uint16
	// Optional header, Specifies the transactional crc64 for the body, to be validated by the service.
	TransactionalContentCRC64 *[]byte
	// Specify the transactional md5 for the body, to be validated by the service.
	TransactionalContentMD5 *[]byte
}

func (o *UploadReaderAtToBlockBlobOption) getStageBlockOptions() *StageBlockOptions {
	leaseAccessConditions, _ := exported.FormatBlobAccessConditions(o.AccessConditions)
	return &StageBlockOptions{
		CpkInfo:               o.CpkInfo,
		CpkScopeInfo:          o.CpkScopeInfo,
		LeaseAccessConditions: leaseAccessConditions,
	}
}

func (o *UploadReaderAtToBlockBlobOption) getUploadBlockBlobOptions() *UploadOptions {
	return &UploadOptions{
		Tags:             o.Tags,
		Metadata:         o.Metadata,
		Tier:             o.AccessTier,
		HTTPHeaders:      o.HTTPHeaders,
		AccessConditions: o.AccessConditions,
		CpkInfo:          o.CpkInfo,
		CpkScopeInfo:     o.CpkScopeInfo,
	}
}

func (o *UploadReaderAtToBlockBlobOption) getCommitBlockListOptions() *CommitBlockListOptions {
	return &CommitBlockListOptions{
		Tags:         o.Tags,
		Metadata:     o.Metadata,
		Tier:         o.AccessTier,
		HTTPHeaders:  o.HTTPHeaders,
		CpkInfo:      o.CpkInfo,
		CpkScopeInfo: o.CpkScopeInfo,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// UploadStreamOptions provides set of configurations for UploadStream operation
type UploadStreamOptions struct {
	// TransferManager provides a TransferManager that controls buffer allocation/reuse and
	// concurrency. This overrides BufferSize and MaxBuffers if set.
	TransferManager      shared.TransferManager
	transferMangerNotSet bool
	// BufferSize sizes the buffer used to read data from source. If < 1 MiB, format to 1 MiB.
	BufferSize int
	// MaxBuffers defines the number of simultaneous uploads will be performed to upload the file.
	MaxBuffers       int
	HTTPHeaders      *blob.HTTPHeaders
	Metadata         map[string]string
	AccessConditions *blob.AccessConditions
	AccessTier       *blob.AccessTier
	Tags             map[string]string
	CpkInfo          *blob.CpkInfo
	CpkScopeInfo     *blob.CpkScopeInfo
}

func (u *UploadStreamOptions) format() error {
	if u == nil || u.TransferManager != nil {
		return nil
	}

	if u.MaxBuffers == 0 {
		u.MaxBuffers = 1
	}

	if u.BufferSize < _1MiB {
		u.BufferSize = _1MiB
	}

	var err error
	u.TransferManager, err = shared.NewStaticBuffer(u.BufferSize, u.MaxBuffers)
	if err != nil {
		return fmt.Errorf("bug: default transfer manager could not be created: %s", err)
	}
	u.transferMangerNotSet = true
	return nil
}

func (u *UploadStreamOptions) getStageBlockOptions() *StageBlockOptions {
	leaseAccessConditions, _ := exported.FormatBlobAccessConditions(u.AccessConditions)
	return &StageBlockOptions{
		CpkInfo:               u.CpkInfo,
		CpkScopeInfo:          u.CpkScopeInfo,
		LeaseAccessConditions: leaseAccessConditions,
	}
}

func (u *UploadStreamOptions) getCommitBlockListOptions() *CommitBlockListOptions {
	options := &CommitBlockListOptions{
		Tags:             u.Tags,
		Metadata:         u.Metadata,
		Tier:             u.AccessTier,
		HTTPHeaders:      u.HTTPHeaders,
		CpkInfo:          u.CpkInfo,
		CpkScopeInfo:     u.CpkScopeInfo,
		AccessConditions: u.AccessConditions,
	}

	return options
}
