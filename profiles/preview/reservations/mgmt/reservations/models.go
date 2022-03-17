//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package reservations

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/reservations/mgmt/2017-11-01/reservations"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AppliedScopeType = original.AppliedScopeType

const (
	Shared AppliedScopeType = original.Shared
	Single AppliedScopeType = original.Single
)

type AppliedScopeType1 = original.AppliedScopeType1

const (
	AppliedScopeType1Shared AppliedScopeType1 = original.AppliedScopeType1Shared
	AppliedScopeType1Single AppliedScopeType1 = original.AppliedScopeType1Single
)

type Code = original.Code

const (
	ActivateQuoteFailed                           Code = original.ActivateQuoteFailed
	AppliedScopesNotAssociatedWithCommerceAccount Code = original.AppliedScopesNotAssociatedWithCommerceAccount
	AppliedScopesSameAsExisting                   Code = original.AppliedScopesSameAsExisting
	AuthorizationFailed                           Code = original.AuthorizationFailed
	BadRequest                                    Code = original.BadRequest
	BillingCustomerInputError                     Code = original.BillingCustomerInputError
	BillingError                                  Code = original.BillingError
	BillingPaymentInstrumentHardError             Code = original.BillingPaymentInstrumentHardError
	BillingPaymentInstrumentSoftError             Code = original.BillingPaymentInstrumentSoftError
	BillingScopeIDCannotBeChanged                 Code = original.BillingScopeIDCannotBeChanged
	BillingTransientError                         Code = original.BillingTransientError
	CalculatePriceFailed                          Code = original.CalculatePriceFailed
	CapacityUpdateScopesFailed                    Code = original.CapacityUpdateScopesFailed
	ClientCertificateThumbprintNotSet             Code = original.ClientCertificateThumbprintNotSet
	CreateQuoteFailed                             Code = original.CreateQuoteFailed
	Forbidden                                     Code = original.Forbidden
	FulfillmentConfigurationError                 Code = original.FulfillmentConfigurationError
	FulfillmentError                              Code = original.FulfillmentError
	FulfillmentOutOfStockError                    Code = original.FulfillmentOutOfStockError
	FulfillmentTransientError                     Code = original.FulfillmentTransientError
	HTTPMethodNotSupported                        Code = original.HTTPMethodNotSupported
	InternalServerError                           Code = original.InternalServerError
	InvalidAccessToken                            Code = original.InvalidAccessToken
	InvalidFulfillmentRequestParameters           Code = original.InvalidFulfillmentRequestParameters
	InvalidHealthCheckType                        Code = original.InvalidHealthCheckType
	InvalidLocationID                             Code = original.InvalidLocationID
	InvalidRefundQuantity                         Code = original.InvalidRefundQuantity
	InvalidRequestContent                         Code = original.InvalidRequestContent
	InvalidRequestURI                             Code = original.InvalidRequestURI
	InvalidReservationID                          Code = original.InvalidReservationID
	InvalidReservationOrderID                     Code = original.InvalidReservationOrderID
	InvalidSingleAppliedScopesCount               Code = original.InvalidSingleAppliedScopesCount
	InvalidSubscriptionID                         Code = original.InvalidSubscriptionID
	InvalidTenantID                               Code = original.InvalidTenantID
	MissingAppliedScopesForSingle                 Code = original.MissingAppliedScopesForSingle
	MissingTenantID                               Code = original.MissingTenantID
	NonsupportedAccountID                         Code = original.NonsupportedAccountID
	NotSpecified                                  Code = original.NotSpecified
	NotSupportedCountry                           Code = original.NotSupportedCountry
	NoValidReservationsToReRate                   Code = original.NoValidReservationsToReRate
	OperationCannotBePerformedInCurrentState      Code = original.OperationCannotBePerformedInCurrentState
	OperationFailed                               Code = original.OperationFailed
	PaymentInstrumentNotFound                     Code = original.PaymentInstrumentNotFound
	PurchaseError                                 Code = original.PurchaseError
	ReRateOnlyAllowedForEA                        Code = original.ReRateOnlyAllowedForEA
	ReservationIDNotInReservationOrder            Code = original.ReservationIDNotInReservationOrder
	ReservationOrderCreationFailed                Code = original.ReservationOrderCreationFailed
	ReservationOrderIDAlreadyExists               Code = original.ReservationOrderIDAlreadyExists
	ReservationOrderNotEnabled                    Code = original.ReservationOrderNotEnabled
	ReservationOrderNotFound                      Code = original.ReservationOrderNotFound
	RiskCheckFailed                               Code = original.RiskCheckFailed
	RoleAssignmentCreationFailed                  Code = original.RoleAssignmentCreationFailed
	ServerTimeout                                 Code = original.ServerTimeout
	UnauthenticatedRequestsThrottled              Code = original.UnauthenticatedRequestsThrottled
	UnsupportedReservationTerm                    Code = original.UnsupportedReservationTerm
)

type Kind = original.Kind

const (
	MicrosoftCompute Kind = original.MicrosoftCompute
)

type Location = original.Location

const (
	Australiaeast      Location = original.Australiaeast
	Australiasoutheast Location = original.Australiasoutheast
	Brazilsouth        Location = original.Brazilsouth
	Canadacentral      Location = original.Canadacentral
	Canadaeast         Location = original.Canadaeast
	Centralindia       Location = original.Centralindia
	Centralus          Location = original.Centralus
	Eastasia           Location = original.Eastasia
	Eastus             Location = original.Eastus
	Eastus2            Location = original.Eastus2
	Japaneast          Location = original.Japaneast
	Japanwest          Location = original.Japanwest
	Northcentralus     Location = original.Northcentralus
	Northeurope        Location = original.Northeurope
	Southcentralus     Location = original.Southcentralus
	Southeastasia      Location = original.Southeastasia
	Southindia         Location = original.Southindia
	Uksouth            Location = original.Uksouth
	Ukwest             Location = original.Ukwest
	Westcentralus      Location = original.Westcentralus
	Westeurope         Location = original.Westeurope
	Westindia          Location = original.Westindia
	Westus             Location = original.Westus
	Westus2            Location = original.Westus2
)

type ProvisioningState = original.ProvisioningState

const (
	BillingFailed         ProvisioningState = original.BillingFailed
	Cancelled             ProvisioningState = original.Cancelled
	ConfirmedBilling      ProvisioningState = original.ConfirmedBilling
	ConfirmedResourceHold ProvisioningState = original.ConfirmedResourceHold
	Created               ProvisioningState = original.Created
	Creating              ProvisioningState = original.Creating
	Expired               ProvisioningState = original.Expired
	Failed                ProvisioningState = original.Failed
	Merged                ProvisioningState = original.Merged
	PendingBilling        ProvisioningState = original.PendingBilling
	PendingResourceHold   ProvisioningState = original.PendingResourceHold
	Split                 ProvisioningState = original.Split
	Succeeded             ProvisioningState = original.Succeeded
)

type ProvisioningState1 = original.ProvisioningState1

const (
	ProvisioningState1BillingFailed         ProvisioningState1 = original.ProvisioningState1BillingFailed
	ProvisioningState1Cancelled             ProvisioningState1 = original.ProvisioningState1Cancelled
	ProvisioningState1ConfirmedBilling      ProvisioningState1 = original.ProvisioningState1ConfirmedBilling
	ProvisioningState1ConfirmedResourceHold ProvisioningState1 = original.ProvisioningState1ConfirmedResourceHold
	ProvisioningState1Created               ProvisioningState1 = original.ProvisioningState1Created
	ProvisioningState1Creating              ProvisioningState1 = original.ProvisioningState1Creating
	ProvisioningState1Expired               ProvisioningState1 = original.ProvisioningState1Expired
	ProvisioningState1Failed                ProvisioningState1 = original.ProvisioningState1Failed
	ProvisioningState1Merged                ProvisioningState1 = original.ProvisioningState1Merged
	ProvisioningState1PendingBilling        ProvisioningState1 = original.ProvisioningState1PendingBilling
	ProvisioningState1PendingResourceHold   ProvisioningState1 = original.ProvisioningState1PendingResourceHold
	ProvisioningState1Split                 ProvisioningState1 = original.ProvisioningState1Split
	ProvisioningState1Succeeded             ProvisioningState1 = original.ProvisioningState1Succeeded
)

type StatusCode = original.StatusCode

const (
	StatusCodeActive                 StatusCode = original.StatusCodeActive
	StatusCodeExpired                StatusCode = original.StatusCodeExpired
	StatusCodeMerged                 StatusCode = original.StatusCodeMerged
	StatusCodeNone                   StatusCode = original.StatusCodeNone
	StatusCodePaymentInstrumentError StatusCode = original.StatusCodePaymentInstrumentError
	StatusCodePending                StatusCode = original.StatusCodePending
	StatusCodePurchaseError          StatusCode = original.StatusCodePurchaseError
	StatusCodeSplit                  StatusCode = original.StatusCodeSplit
	StatusCodeSucceeded              StatusCode = original.StatusCodeSucceeded
)

type Term = original.Term

const (
	P1Y Term = original.P1Y
	P3Y Term = original.P3Y
)

type AppliedReservationList = original.AppliedReservationList
type AppliedReservations = original.AppliedReservations
type AppliedReservationsProperties = original.AppliedReservationsProperties
type BaseClient = original.BaseClient
type Catalog = original.Catalog
type Client = original.Client
type Error = original.Error
type ExtendedErrorInfo = original.ExtendedErrorInfo
type ExtendedStatusInfo = original.ExtendedStatusInfo
type List = original.List
type ListCatalog = original.ListCatalog
type ListIterator = original.ListIterator
type ListPage = original.ListPage
type ListResponse = original.ListResponse
type MergeProperties = original.MergeProperties
type MergePropertiesType = original.MergePropertiesType
type MergeRequest = original.MergeRequest
type OperationClient = original.OperationClient
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationResponse = original.OperationResponse
type OrderClient = original.OrderClient
type OrderList = original.OrderList
type OrderListIterator = original.OrderListIterator
type OrderListPage = original.OrderListPage
type OrderProperties = original.OrderProperties
type OrderResponse = original.OrderResponse
type Patch = original.Patch
type PatchProperties = original.PatchProperties
type Properties = original.Properties
type ReservationMergeFuture = original.ReservationMergeFuture
type ReservationUpdateFuture = original.ReservationUpdateFuture
type Response = original.Response
type SkuCapability = original.SkuCapability
type SkuName = original.SkuName
type SkuRestriction = original.SkuRestriction
type SplitFuture = original.SplitFuture
type SplitProperties = original.SplitProperties
type SplitPropertiesType = original.SplitPropertiesType
type SplitRequest = original.SplitRequest

func New() BaseClient {
	return original.New()
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewListIterator(page ListPage) ListIterator {
	return original.NewListIterator(page)
}
func NewListPage(cur List, getNextPage func(context.Context, List) (List, error)) ListPage {
	return original.NewListPage(cur, getNextPage)
}
func NewOperationClient() OperationClient {
	return original.NewOperationClient()
}
func NewOperationClientWithBaseURI(baseURI string) OperationClient {
	return original.NewOperationClientWithBaseURI(baseURI)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOrderClient() OrderClient {
	return original.NewOrderClient()
}
func NewOrderClientWithBaseURI(baseURI string) OrderClient {
	return original.NewOrderClientWithBaseURI(baseURI)
}
func NewOrderListIterator(page OrderListPage) OrderListIterator {
	return original.NewOrderListIterator(page)
}
func NewOrderListPage(cur OrderList, getNextPage func(context.Context, OrderList) (OrderList, error)) OrderListPage {
	return original.NewOrderListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleAppliedScopeType1Values() []AppliedScopeType1 {
	return original.PossibleAppliedScopeType1Values()
}
func PossibleAppliedScopeTypeValues() []AppliedScopeType {
	return original.PossibleAppliedScopeTypeValues()
}
func PossibleCodeValues() []Code {
	return original.PossibleCodeValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLocationValues() []Location {
	return original.PossibleLocationValues()
}
func PossibleProvisioningState1Values() []ProvisioningState1 {
	return original.PossibleProvisioningState1Values()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleStatusCodeValues() []StatusCode {
	return original.PossibleStatusCodeValues()
}
func PossibleTermValues() []Term {
	return original.PossibleTermValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
