//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package customerlockbox

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/preview/customerlockbox/mgmt/2018-02-28-preview/customerlockbox"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Decision = original.Decision

const (
	Approve Decision = original.Approve
	Deny    Decision = original.Deny
)

type Status = original.Status

const (
	Approved     Status = original.Approved
	Approving    Status = original.Approving
	Completed    Status = original.Completed
	Completing   Status = original.Completing
	Denied       Status = original.Denied
	Denying      Status = original.Denying
	Error        Status = original.Error
	Expired      Status = original.Expired
	Initializing Status = original.Initializing
	Pending      Status = original.Pending
	Revoked      Status = original.Revoked
	Revoking     Status = original.Revoking
	Unknown      Status = original.Unknown
)

type Approval = original.Approval
type BaseClient = original.BaseClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorAdditionalInfoInfo = original.ErrorAdditionalInfoInfo
type ErrorBody = original.ErrorBody
type ErrorResponse = original.ErrorResponse
type LockboxRequestResponse = original.LockboxRequestResponse
type LockboxRequestResponseProperties = original.LockboxRequestResponseProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type RequestListResult = original.RequestListResult
type RequestListResultIterator = original.RequestListResultIterator
type RequestListResultPage = original.RequestListResultPage
type RequestsClient = original.RequestsClient

func New() BaseClient {
	return original.New()
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewRequestListResultIterator(page RequestListResultPage) RequestListResultIterator {
	return original.NewRequestListResultIterator(page)
}
func NewRequestListResultPage(cur RequestListResult, getNextPage func(context.Context, RequestListResult) (RequestListResult, error)) RequestListResultPage {
	return original.NewRequestListResultPage(cur, getNextPage)
}
func NewRequestsClient() RequestsClient {
	return original.NewRequestsClient()
}
func NewRequestsClientWithBaseURI(baseURI string) RequestsClient {
	return original.NewRequestsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleDecisionValues() []Decision {
	return original.PossibleDecisionValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
