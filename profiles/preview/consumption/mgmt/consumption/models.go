// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package consumption

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2019-10-01/consumption"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BillingFrequency = original.BillingFrequency

const (
	BillingFrequencyMonth   BillingFrequency = original.BillingFrequencyMonth
	BillingFrequencyQuarter BillingFrequency = original.BillingFrequencyQuarter
	BillingFrequencyYear    BillingFrequency = original.BillingFrequencyYear
)

type Bound = original.Bound

const (
	BoundLower Bound = original.BoundLower
	BoundUpper Bound = original.BoundUpper
)

type ChargeType = original.ChargeType

const (
	ChargeTypeActual   ChargeType = original.ChargeTypeActual
	ChargeTypeForecast ChargeType = original.ChargeTypeForecast
)

type CultureCode = original.CultureCode

const (
	CultureCodeCsCz CultureCode = original.CultureCodeCsCz
	CultureCodeDaDk CultureCode = original.CultureCodeDaDk
	CultureCodeDeDe CultureCode = original.CultureCodeDeDe
	CultureCodeEnGb CultureCode = original.CultureCodeEnGb
	CultureCodeEnUs CultureCode = original.CultureCodeEnUs
	CultureCodeEsEs CultureCode = original.CultureCodeEsEs
	CultureCodeFrFr CultureCode = original.CultureCodeFrFr
	CultureCodeHuHu CultureCode = original.CultureCodeHuHu
	CultureCodeItIt CultureCode = original.CultureCodeItIt
	CultureCodeJaJp CultureCode = original.CultureCodeJaJp
	CultureCodeKoKr CultureCode = original.CultureCodeKoKr
	CultureCodeNbNo CultureCode = original.CultureCodeNbNo
	CultureCodeNlNl CultureCode = original.CultureCodeNlNl
	CultureCodePlPl CultureCode = original.CultureCodePlPl
	CultureCodePtBr CultureCode = original.CultureCodePtBr
	CultureCodePtPt CultureCode = original.CultureCodePtPt
	CultureCodeRuRu CultureCode = original.CultureCodeRuRu
	CultureCodeSvSe CultureCode = original.CultureCodeSvSe
	CultureCodeTrTr CultureCode = original.CultureCodeTrTr
	CultureCodeZhCn CultureCode = original.CultureCodeZhCn
	CultureCodeZhTw CultureCode = original.CultureCodeZhTw
)

type Datagrain = original.Datagrain

const (
	DatagrainDailyGrain   Datagrain = original.DatagrainDailyGrain
	DatagrainMonthlyGrain Datagrain = original.DatagrainMonthlyGrain
)

type EventType = original.EventType

const (
	EventTypeNewCredit            EventType = original.EventTypeNewCredit
	EventTypePendingAdjustments   EventType = original.EventTypePendingAdjustments
	EventTypePendingCharges       EventType = original.EventTypePendingCharges
	EventTypePendingExpiredCredit EventType = original.EventTypePendingExpiredCredit
	EventTypePendingNewCredit     EventType = original.EventTypePendingNewCredit
	EventTypeSettledCharges       EventType = original.EventTypeSettledCharges
	EventTypeUnKnown              EventType = original.EventTypeUnKnown
)

type Grain = original.Grain

const (
	GrainDaily   Grain = original.GrainDaily
	GrainMonthly Grain = original.GrainMonthly
	GrainYearly  Grain = original.GrainYearly
)

type Kind = original.Kind

const (
	KindLegacy      Kind = original.KindLegacy
	KindModern      Kind = original.KindModern
	KindUsageDetail Kind = original.KindUsageDetail
)

type KindBasicChargeSummary = original.KindBasicChargeSummary

const (
	KindBasicChargeSummaryKindChargeSummary KindBasicChargeSummary = original.KindBasicChargeSummaryKindChargeSummary
	KindBasicChargeSummaryKindLegacy        KindBasicChargeSummary = original.KindBasicChargeSummaryKindLegacy
	KindBasicChargeSummaryKindModern        KindBasicChargeSummary = original.KindBasicChargeSummaryKindModern
)

type KindBasicReservationRecommendation = original.KindBasicReservationRecommendation

const (
	KindBasicReservationRecommendationKindLegacy                    KindBasicReservationRecommendation = original.KindBasicReservationRecommendationKindLegacy
	KindBasicReservationRecommendationKindModern                    KindBasicReservationRecommendation = original.KindBasicReservationRecommendationKindModern
	KindBasicReservationRecommendationKindReservationRecommendation KindBasicReservationRecommendation = original.KindBasicReservationRecommendationKindReservationRecommendation
)

type LookBackPeriod = original.LookBackPeriod

const (
	LookBackPeriodLast07Days LookBackPeriod = original.LookBackPeriodLast07Days
	LookBackPeriodLast30Days LookBackPeriod = original.LookBackPeriodLast30Days
	LookBackPeriodLast60Days LookBackPeriod = original.LookBackPeriodLast60Days
)

type LotSource = original.LotSource

const (
	LotSourcePromotionalCredit LotSource = original.LotSourcePromotionalCredit
	LotSourcePurchasedCredit   LotSource = original.LotSourcePurchasedCredit
)

type Metrictype = original.Metrictype

const (
	MetrictypeActualCostMetricType    Metrictype = original.MetrictypeActualCostMetricType
	MetrictypeAmortizedCostMetricType Metrictype = original.MetrictypeAmortizedCostMetricType
	MetrictypeUsageMetricType         Metrictype = original.MetrictypeUsageMetricType
)

type OperatorType = original.OperatorType

const (
	OperatorTypeEqualTo              OperatorType = original.OperatorTypeEqualTo
	OperatorTypeGreaterThan          OperatorType = original.OperatorTypeGreaterThan
	OperatorTypeGreaterThanOrEqualTo OperatorType = original.OperatorTypeGreaterThanOrEqualTo
)

type Scope12 = original.Scope12

const (
	Scope12Shared Scope12 = original.Scope12Shared
	Scope12Single Scope12 = original.Scope12Single
)

type Scope14 = original.Scope14

const (
	Scope14Shared Scope14 = original.Scope14Shared
	Scope14Single Scope14 = original.Scope14Single
)

type Term = original.Term

const (
	TermP1Y Term = original.TermP1Y
	TermP3Y Term = original.TermP3Y
)

type ThresholdType = original.ThresholdType

const (
	ThresholdTypeActual ThresholdType = original.ThresholdTypeActual
)

type TimeGrainType = original.TimeGrainType

const (
	TimeGrainTypeAnnually       TimeGrainType = original.TimeGrainTypeAnnually
	TimeGrainTypeBillingAnnual  TimeGrainType = original.TimeGrainTypeBillingAnnual
	TimeGrainTypeBillingMonth   TimeGrainType = original.TimeGrainTypeBillingMonth
	TimeGrainTypeBillingQuarter TimeGrainType = original.TimeGrainTypeBillingQuarter
	TimeGrainTypeMonthly        TimeGrainType = original.TimeGrainTypeMonthly
	TimeGrainTypeQuarterly      TimeGrainType = original.TimeGrainTypeQuarterly
)

type AggregatedCostClient = original.AggregatedCostClient
type Amount = original.Amount
type AmountWithExchangeRate = original.AmountWithExchangeRate
type Balance = original.Balance
type BalanceProperties = original.BalanceProperties
type BalancePropertiesAdjustmentDetailsItem = original.BalancePropertiesAdjustmentDetailsItem
type BalancePropertiesNewPurchasesDetailsItem = original.BalancePropertiesNewPurchasesDetailsItem
type BalancesClient = original.BalancesClient
type BaseClient = original.BaseClient
type BasicChargeSummary = original.BasicChargeSummary
type BasicReservationRecommendation = original.BasicReservationRecommendation
type BasicUsageDetail = original.BasicUsageDetail
type Budget = original.Budget
type BudgetComparisonExpression = original.BudgetComparisonExpression
type BudgetFilter = original.BudgetFilter
type BudgetFilterProperties = original.BudgetFilterProperties
type BudgetProperties = original.BudgetProperties
type BudgetTimePeriod = original.BudgetTimePeriod
type BudgetsClient = original.BudgetsClient
type BudgetsListResult = original.BudgetsListResult
type BudgetsListResultIterator = original.BudgetsListResultIterator
type BudgetsListResultPage = original.BudgetsListResultPage
type ChargeSummary = original.ChargeSummary
type ChargesClient = original.ChargesClient
type ChargesListResult = original.ChargesListResult
type CreditBalanceSummary = original.CreditBalanceSummary
type CreditSummary = original.CreditSummary
type CreditSummaryProperties = original.CreditSummaryProperties
type CreditsClient = original.CreditsClient
type CurrentSpend = original.CurrentSpend
type DownloadProperties = original.DownloadProperties
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type EventProperties = original.EventProperties
type EventSummary = original.EventSummary
type Events = original.Events
type EventsClient = original.EventsClient
type EventsIterator = original.EventsIterator
type EventsPage = original.EventsPage
type Forecast = original.Forecast
type ForecastProperties = original.ForecastProperties
type ForecastPropertiesConfidenceLevelsItem = original.ForecastPropertiesConfidenceLevelsItem
type ForecastSpend = original.ForecastSpend
type ForecastsClient = original.ForecastsClient
type ForecastsListResult = original.ForecastsListResult
type HighCasedErrorDetails = original.HighCasedErrorDetails
type HighCasedErrorResponse = original.HighCasedErrorResponse
type LegacyChargeSummary = original.LegacyChargeSummary
type LegacyChargeSummaryProperties = original.LegacyChargeSummaryProperties
type LegacyReservationRecommendation = original.LegacyReservationRecommendation
type LegacyReservationRecommendationProperties = original.LegacyReservationRecommendationProperties
type LegacyReservationTransaction = original.LegacyReservationTransaction
type LegacyReservationTransactionProperties = original.LegacyReservationTransactionProperties
type LegacyUsageDetail = original.LegacyUsageDetail
type LegacyUsageDetailProperties = original.LegacyUsageDetailProperties
type LotProperties = original.LotProperties
type LotSummary = original.LotSummary
type Lots = original.Lots
type LotsClient = original.LotsClient
type LotsIterator = original.LotsIterator
type LotsPage = original.LotsPage
type ManagementGroupAggregatedCostProperties = original.ManagementGroupAggregatedCostProperties
type ManagementGroupAggregatedCostResult = original.ManagementGroupAggregatedCostResult
type Marketplace = original.Marketplace
type MarketplaceProperties = original.MarketplaceProperties
type MarketplacesClient = original.MarketplacesClient
type MarketplacesListResult = original.MarketplacesListResult
type MarketplacesListResultIterator = original.MarketplacesListResultIterator
type MarketplacesListResultPage = original.MarketplacesListResultPage
type MeterDetails = original.MeterDetails
type MeterDetailsResponse = original.MeterDetailsResponse
type ModernChargeSummary = original.ModernChargeSummary
type ModernChargeSummaryProperties = original.ModernChargeSummaryProperties
type ModernReservationRecommendation = original.ModernReservationRecommendation
type ModernReservationRecommendationProperties = original.ModernReservationRecommendationProperties
type ModernReservationTransaction = original.ModernReservationTransaction
type ModernReservationTransactionProperties = original.ModernReservationTransactionProperties
type ModernReservationTransactionsListResult = original.ModernReservationTransactionsListResult
type ModernReservationTransactionsListResultIterator = original.ModernReservationTransactionsListResultIterator
type ModernReservationTransactionsListResultPage = original.ModernReservationTransactionsListResultPage
type ModernUsageDetail = original.ModernUsageDetail
type ModernUsageDetailProperties = original.ModernUsageDetailProperties
type Notification = original.Notification
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PriceSheetClient = original.PriceSheetClient
type PriceSheetModel = original.PriceSheetModel
type PriceSheetProperties = original.PriceSheetProperties
type PriceSheetResult = original.PriceSheetResult
type ProxyResource = original.ProxyResource
type Reseller = original.Reseller
type ReservationDetail = original.ReservationDetail
type ReservationDetailProperties = original.ReservationDetailProperties
type ReservationDetailsListResult = original.ReservationDetailsListResult
type ReservationDetailsListResultIterator = original.ReservationDetailsListResultIterator
type ReservationDetailsListResultPage = original.ReservationDetailsListResultPage
type ReservationRecommendation = original.ReservationRecommendation
type ReservationRecommendationDetailsCalculatedSavingsProperties = original.ReservationRecommendationDetailsCalculatedSavingsProperties
type ReservationRecommendationDetailsClient = original.ReservationRecommendationDetailsClient
type ReservationRecommendationDetailsModel = original.ReservationRecommendationDetailsModel
type ReservationRecommendationDetailsProperties = original.ReservationRecommendationDetailsProperties
type ReservationRecommendationDetailsResourceProperties = original.ReservationRecommendationDetailsResourceProperties
type ReservationRecommendationDetailsSavingsProperties = original.ReservationRecommendationDetailsSavingsProperties
type ReservationRecommendationDetailsUsageProperties = original.ReservationRecommendationDetailsUsageProperties
type ReservationRecommendationsClient = original.ReservationRecommendationsClient
type ReservationRecommendationsListResult = original.ReservationRecommendationsListResult
type ReservationRecommendationsListResultIterator = original.ReservationRecommendationsListResultIterator
type ReservationRecommendationsListResultPage = original.ReservationRecommendationsListResultPage
type ReservationSummariesListResult = original.ReservationSummariesListResult
type ReservationSummariesListResultIterator = original.ReservationSummariesListResultIterator
type ReservationSummariesListResultPage = original.ReservationSummariesListResultPage
type ReservationSummary = original.ReservationSummary
type ReservationSummaryProperties = original.ReservationSummaryProperties
type ReservationTransaction = original.ReservationTransaction
type ReservationTransactionResource = original.ReservationTransactionResource
type ReservationTransactionsClient = original.ReservationTransactionsClient
type ReservationTransactionsListResult = original.ReservationTransactionsListResult
type ReservationTransactionsListResultIterator = original.ReservationTransactionsListResultIterator
type ReservationTransactionsListResultPage = original.ReservationTransactionsListResultPage
type ReservationsDetailsClient = original.ReservationsDetailsClient
type ReservationsSummariesClient = original.ReservationsSummariesClient
type Resource = original.Resource
type ResourceAttributes = original.ResourceAttributes
type SkuProperty = original.SkuProperty
type Tag = original.Tag
type TagProperties = original.TagProperties
type TagsClient = original.TagsClient
type TagsResult = original.TagsResult
type UsageDetail = original.UsageDetail
type UsageDetailsClient = original.UsageDetailsClient
type UsageDetailsListResult = original.UsageDetailsListResult
type UsageDetailsListResultIterator = original.UsageDetailsListResultIterator
type UsageDetailsListResultPage = original.UsageDetailsListResultPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAggregatedCostClient(subscriptionID string) AggregatedCostClient {
	return original.NewAggregatedCostClient(subscriptionID)
}
func NewAggregatedCostClientWithBaseURI(baseURI string, subscriptionID string) AggregatedCostClient {
	return original.NewAggregatedCostClientWithBaseURI(baseURI, subscriptionID)
}
func NewBalancesClient(subscriptionID string) BalancesClient {
	return original.NewBalancesClient(subscriptionID)
}
func NewBalancesClientWithBaseURI(baseURI string, subscriptionID string) BalancesClient {
	return original.NewBalancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewBudgetsClient(subscriptionID string) BudgetsClient {
	return original.NewBudgetsClient(subscriptionID)
}
func NewBudgetsClientWithBaseURI(baseURI string, subscriptionID string) BudgetsClient {
	return original.NewBudgetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBudgetsListResultIterator(page BudgetsListResultPage) BudgetsListResultIterator {
	return original.NewBudgetsListResultIterator(page)
}
func NewBudgetsListResultPage(cur BudgetsListResult, getNextPage func(context.Context, BudgetsListResult) (BudgetsListResult, error)) BudgetsListResultPage {
	return original.NewBudgetsListResultPage(cur, getNextPage)
}
func NewChargesClient(subscriptionID string) ChargesClient {
	return original.NewChargesClient(subscriptionID)
}
func NewChargesClientWithBaseURI(baseURI string, subscriptionID string) ChargesClient {
	return original.NewChargesClientWithBaseURI(baseURI, subscriptionID)
}
func NewCreditsClient(subscriptionID string) CreditsClient {
	return original.NewCreditsClient(subscriptionID)
}
func NewCreditsClientWithBaseURI(baseURI string, subscriptionID string) CreditsClient {
	return original.NewCreditsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventsClient(subscriptionID string) EventsClient {
	return original.NewEventsClient(subscriptionID)
}
func NewEventsClientWithBaseURI(baseURI string, subscriptionID string) EventsClient {
	return original.NewEventsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventsIterator(page EventsPage) EventsIterator {
	return original.NewEventsIterator(page)
}
func NewEventsPage(cur Events, getNextPage func(context.Context, Events) (Events, error)) EventsPage {
	return original.NewEventsPage(cur, getNextPage)
}
func NewForecastsClient(subscriptionID string) ForecastsClient {
	return original.NewForecastsClient(subscriptionID)
}
func NewForecastsClientWithBaseURI(baseURI string, subscriptionID string) ForecastsClient {
	return original.NewForecastsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLotsClient(subscriptionID string) LotsClient {
	return original.NewLotsClient(subscriptionID)
}
func NewLotsClientWithBaseURI(baseURI string, subscriptionID string) LotsClient {
	return original.NewLotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLotsIterator(page LotsPage) LotsIterator {
	return original.NewLotsIterator(page)
}
func NewLotsPage(cur Lots, getNextPage func(context.Context, Lots) (Lots, error)) LotsPage {
	return original.NewLotsPage(cur, getNextPage)
}
func NewMarketplacesClient(subscriptionID string) MarketplacesClient {
	return original.NewMarketplacesClient(subscriptionID)
}
func NewMarketplacesClientWithBaseURI(baseURI string, subscriptionID string) MarketplacesClient {
	return original.NewMarketplacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMarketplacesListResultIterator(page MarketplacesListResultPage) MarketplacesListResultIterator {
	return original.NewMarketplacesListResultIterator(page)
}
func NewMarketplacesListResultPage(cur MarketplacesListResult, getNextPage func(context.Context, MarketplacesListResult) (MarketplacesListResult, error)) MarketplacesListResultPage {
	return original.NewMarketplacesListResultPage(cur, getNextPage)
}
func NewModernReservationTransactionsListResultIterator(page ModernReservationTransactionsListResultPage) ModernReservationTransactionsListResultIterator {
	return original.NewModernReservationTransactionsListResultIterator(page)
}
func NewModernReservationTransactionsListResultPage(cur ModernReservationTransactionsListResult, getNextPage func(context.Context, ModernReservationTransactionsListResult) (ModernReservationTransactionsListResult, error)) ModernReservationTransactionsListResultPage {
	return original.NewModernReservationTransactionsListResultPage(cur, getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPriceSheetClient(subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClient(subscriptionID)
}
func NewPriceSheetClientWithBaseURI(baseURI string, subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationDetailsListResultIterator(page ReservationDetailsListResultPage) ReservationDetailsListResultIterator {
	return original.NewReservationDetailsListResultIterator(page)
}
func NewReservationDetailsListResultPage(cur ReservationDetailsListResult, getNextPage func(context.Context, ReservationDetailsListResult) (ReservationDetailsListResult, error)) ReservationDetailsListResultPage {
	return original.NewReservationDetailsListResultPage(cur, getNextPage)
}
func NewReservationRecommendationDetailsClient(subscriptionID string) ReservationRecommendationDetailsClient {
	return original.NewReservationRecommendationDetailsClient(subscriptionID)
}
func NewReservationRecommendationDetailsClientWithBaseURI(baseURI string, subscriptionID string) ReservationRecommendationDetailsClient {
	return original.NewReservationRecommendationDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationRecommendationsClient(subscriptionID string) ReservationRecommendationsClient {
	return original.NewReservationRecommendationsClient(subscriptionID)
}
func NewReservationRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) ReservationRecommendationsClient {
	return original.NewReservationRecommendationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationRecommendationsListResultIterator(page ReservationRecommendationsListResultPage) ReservationRecommendationsListResultIterator {
	return original.NewReservationRecommendationsListResultIterator(page)
}
func NewReservationRecommendationsListResultPage(cur ReservationRecommendationsListResult, getNextPage func(context.Context, ReservationRecommendationsListResult) (ReservationRecommendationsListResult, error)) ReservationRecommendationsListResultPage {
	return original.NewReservationRecommendationsListResultPage(cur, getNextPage)
}
func NewReservationSummariesListResultIterator(page ReservationSummariesListResultPage) ReservationSummariesListResultIterator {
	return original.NewReservationSummariesListResultIterator(page)
}
func NewReservationSummariesListResultPage(cur ReservationSummariesListResult, getNextPage func(context.Context, ReservationSummariesListResult) (ReservationSummariesListResult, error)) ReservationSummariesListResultPage {
	return original.NewReservationSummariesListResultPage(cur, getNextPage)
}
func NewReservationTransactionsClient(subscriptionID string) ReservationTransactionsClient {
	return original.NewReservationTransactionsClient(subscriptionID)
}
func NewReservationTransactionsClientWithBaseURI(baseURI string, subscriptionID string) ReservationTransactionsClient {
	return original.NewReservationTransactionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationTransactionsListResultIterator(page ReservationTransactionsListResultPage) ReservationTransactionsListResultIterator {
	return original.NewReservationTransactionsListResultIterator(page)
}
func NewReservationTransactionsListResultPage(cur ReservationTransactionsListResult, getNextPage func(context.Context, ReservationTransactionsListResult) (ReservationTransactionsListResult, error)) ReservationTransactionsListResultPage {
	return original.NewReservationTransactionsListResultPage(cur, getNextPage)
}
func NewReservationsDetailsClient(subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClient(subscriptionID)
}
func NewReservationsDetailsClientWithBaseURI(baseURI string, subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsSummariesClient(subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClient(subscriptionID)
}
func NewReservationsSummariesClientWithBaseURI(baseURI string, subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsClient(subscriptionID string) TagsClient {
	return original.NewTagsClient(subscriptionID)
}
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return original.NewTagsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageDetailsClient(subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClient(subscriptionID)
}
func NewUsageDetailsClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageDetailsListResultIterator(page UsageDetailsListResultPage) UsageDetailsListResultIterator {
	return original.NewUsageDetailsListResultIterator(page)
}
func NewUsageDetailsListResultPage(cur UsageDetailsListResult, getNextPage func(context.Context, UsageDetailsListResult) (UsageDetailsListResult, error)) UsageDetailsListResultPage {
	return original.NewUsageDetailsListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleBillingFrequencyValues() []BillingFrequency {
	return original.PossibleBillingFrequencyValues()
}
func PossibleBoundValues() []Bound {
	return original.PossibleBoundValues()
}
func PossibleChargeTypeValues() []ChargeType {
	return original.PossibleChargeTypeValues()
}
func PossibleCultureCodeValues() []CultureCode {
	return original.PossibleCultureCodeValues()
}
func PossibleDatagrainValues() []Datagrain {
	return original.PossibleDatagrainValues()
}
func PossibleEventTypeValues() []EventType {
	return original.PossibleEventTypeValues()
}
func PossibleGrainValues() []Grain {
	return original.PossibleGrainValues()
}
func PossibleKindBasicChargeSummaryValues() []KindBasicChargeSummary {
	return original.PossibleKindBasicChargeSummaryValues()
}
func PossibleKindBasicReservationRecommendationValues() []KindBasicReservationRecommendation {
	return original.PossibleKindBasicReservationRecommendationValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLookBackPeriodValues() []LookBackPeriod {
	return original.PossibleLookBackPeriodValues()
}
func PossibleLotSourceValues() []LotSource {
	return original.PossibleLotSourceValues()
}
func PossibleMetrictypeValues() []Metrictype {
	return original.PossibleMetrictypeValues()
}
func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}
func PossibleScope12Values() []Scope12 {
	return original.PossibleScope12Values()
}
func PossibleScope14Values() []Scope14 {
	return original.PossibleScope14Values()
}
func PossibleTermValues() []Term {
	return original.PossibleTermValues()
}
func PossibleThresholdTypeValues() []ThresholdType {
	return original.PossibleThresholdTypeValues()
}
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return original.PossibleTimeGrainTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
