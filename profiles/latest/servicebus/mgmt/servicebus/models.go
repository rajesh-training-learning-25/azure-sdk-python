//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package servicebus

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessRights = original.AccessRights

const (
	Listen        AccessRights = original.Listen
	Manage        AccessRights = original.Manage
	SendEnumValue AccessRights = original.SendEnumValue
)

type DefaultAction = original.DefaultAction

const (
	Allow DefaultAction = original.Allow
	Deny  DefaultAction = original.Deny
)

type EncodingCaptureDescription = original.EncodingCaptureDescription

const (
	Avro        EncodingCaptureDescription = original.Avro
	AvroDeflate EncodingCaptureDescription = original.AvroDeflate
)

type EntityStatus = original.EntityStatus

const (
	Active          EntityStatus = original.Active
	Creating        EntityStatus = original.Creating
	Deleting        EntityStatus = original.Deleting
	Disabled        EntityStatus = original.Disabled
	ReceiveDisabled EntityStatus = original.ReceiveDisabled
	Renaming        EntityStatus = original.Renaming
	Restoring       EntityStatus = original.Restoring
	SendDisabled    EntityStatus = original.SendDisabled
	Unknown         EntityStatus = original.Unknown
)

type FilterType = original.FilterType

const (
	FilterTypeCorrelationFilter FilterType = original.FilterTypeCorrelationFilter
	FilterTypeSQLFilter         FilterType = original.FilterTypeSQLFilter
)

type KeyType = original.KeyType

const (
	PrimaryKey   KeyType = original.PrimaryKey
	SecondaryKey KeyType = original.SecondaryKey
)

type NameSpaceType = original.NameSpaceType

const (
	EventHub        NameSpaceType = original.EventHub
	Messaging       NameSpaceType = original.Messaging
	Mixed           NameSpaceType = original.Mixed
	NotificationHub NameSpaceType = original.NotificationHub
	Relay           NameSpaceType = original.Relay
)

type NetworkRuleIPAction = original.NetworkRuleIPAction

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = original.NetworkRuleIPActionAllow
)

type ProvisioningStateDR = original.ProvisioningStateDR

const (
	Accepted  ProvisioningStateDR = original.Accepted
	Failed    ProvisioningStateDR = original.Failed
	Succeeded ProvisioningStateDR = original.Succeeded
)

type RoleDisasterRecovery = original.RoleDisasterRecovery

const (
	Primary               RoleDisasterRecovery = original.Primary
	PrimaryNotReplicating RoleDisasterRecovery = original.PrimaryNotReplicating
	Secondary             RoleDisasterRecovery = original.Secondary
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type UnavailableReason = original.UnavailableReason

const (
	InvalidName                           UnavailableReason = original.InvalidName
	NameInLockdown                        UnavailableReason = original.NameInLockdown
	NameInUse                             UnavailableReason = original.NameInUse
	None                                  UnavailableReason = original.None
	SubscriptionIsDisabled                UnavailableReason = original.SubscriptionIsDisabled
	TooManyNamespaceInCurrentSubscription UnavailableReason = original.TooManyNamespaceInCurrentSubscription
)

type AccessKeys = original.AccessKeys
type Action = original.Action
type ArmDisasterRecovery = original.ArmDisasterRecovery
type ArmDisasterRecoveryListResult = original.ArmDisasterRecoveryListResult
type ArmDisasterRecoveryListResultIterator = original.ArmDisasterRecoveryListResultIterator
type ArmDisasterRecoveryListResultPage = original.ArmDisasterRecoveryListResultPage
type ArmDisasterRecoveryProperties = original.ArmDisasterRecoveryProperties
type BaseClient = original.BaseClient
type CaptureDescription = original.CaptureDescription
type CheckNameAvailability = original.CheckNameAvailability
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CorrelationFilter = original.CorrelationFilter
type Destination = original.Destination
type DestinationProperties = original.DestinationProperties
type DisasterRecoveryConfigsClient = original.DisasterRecoveryConfigsClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type EventHubListResult = original.EventHubListResult
type EventHubListResultIterator = original.EventHubListResultIterator
type EventHubListResultPage = original.EventHubListResultPage
type EventHubsClient = original.EventHubsClient
type Eventhub = original.Eventhub
type EventhubProperties = original.EventhubProperties
type MessageCountDetails = original.MessageCountDetails
type MigrationConfigListResult = original.MigrationConfigListResult
type MigrationConfigListResultIterator = original.MigrationConfigListResultIterator
type MigrationConfigListResultPage = original.MigrationConfigListResultPage
type MigrationConfigProperties = original.MigrationConfigProperties
type MigrationConfigPropertiesProperties = original.MigrationConfigPropertiesProperties
type MigrationConfigsClient = original.MigrationConfigsClient
type MigrationConfigsCreateAndStartMigrationFuture = original.MigrationConfigsCreateAndStartMigrationFuture
type NWRuleSetIPRules = original.NWRuleSetIPRules
type NWRuleSetVirtualNetworkRules = original.NWRuleSetVirtualNetworkRules
type NamespacesClient = original.NamespacesClient
type NamespacesCreateOrUpdateFuture = original.NamespacesCreateOrUpdateFuture
type NamespacesDeleteFuture = original.NamespacesDeleteFuture
type NetworkRuleSet = original.NetworkRuleSet
type NetworkRuleSetListResult = original.NetworkRuleSetListResult
type NetworkRuleSetListResultIterator = original.NetworkRuleSetListResultIterator
type NetworkRuleSetListResultPage = original.NetworkRuleSetListResultPage
type NetworkRuleSetProperties = original.NetworkRuleSetProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PremiumMessagingRegions = original.PremiumMessagingRegions
type PremiumMessagingRegionsClient = original.PremiumMessagingRegionsClient
type PremiumMessagingRegionsListResult = original.PremiumMessagingRegionsListResult
type PremiumMessagingRegionsListResultIterator = original.PremiumMessagingRegionsListResultIterator
type PremiumMessagingRegionsListResultPage = original.PremiumMessagingRegionsListResultPage
type PremiumMessagingRegionsProperties = original.PremiumMessagingRegionsProperties
type QueuesClient = original.QueuesClient
type RegenerateAccessKeyParameters = original.RegenerateAccessKeyParameters
type RegionsClient = original.RegionsClient
type Resource = original.Resource
type ResourceNamespacePatch = original.ResourceNamespacePatch
type Rule = original.Rule
type RuleListResult = original.RuleListResult
type RuleListResultIterator = original.RuleListResultIterator
type RuleListResultPage = original.RuleListResultPage
type Ruleproperties = original.Ruleproperties
type RulesClient = original.RulesClient
type SBAuthorizationRule = original.SBAuthorizationRule
type SBAuthorizationRuleListResult = original.SBAuthorizationRuleListResult
type SBAuthorizationRuleListResultIterator = original.SBAuthorizationRuleListResultIterator
type SBAuthorizationRuleListResultPage = original.SBAuthorizationRuleListResultPage
type SBAuthorizationRuleProperties = original.SBAuthorizationRuleProperties
type SBNamespace = original.SBNamespace
type SBNamespaceListResult = original.SBNamespaceListResult
type SBNamespaceListResultIterator = original.SBNamespaceListResultIterator
type SBNamespaceListResultPage = original.SBNamespaceListResultPage
type SBNamespaceMigrate = original.SBNamespaceMigrate
type SBNamespaceProperties = original.SBNamespaceProperties
type SBNamespaceUpdateParameters = original.SBNamespaceUpdateParameters
type SBQueue = original.SBQueue
type SBQueueListResult = original.SBQueueListResult
type SBQueueListResultIterator = original.SBQueueListResultIterator
type SBQueueListResultPage = original.SBQueueListResultPage
type SBQueueProperties = original.SBQueueProperties
type SBSku = original.SBSku
type SBSubscription = original.SBSubscription
type SBSubscriptionListResult = original.SBSubscriptionListResult
type SBSubscriptionListResultIterator = original.SBSubscriptionListResultIterator
type SBSubscriptionListResultPage = original.SBSubscriptionListResultPage
type SBSubscriptionProperties = original.SBSubscriptionProperties
type SBTopic = original.SBTopic
type SBTopicListResult = original.SBTopicListResult
type SBTopicListResultIterator = original.SBTopicListResultIterator
type SBTopicListResultPage = original.SBTopicListResultPage
type SBTopicProperties = original.SBTopicProperties
type SQLFilter = original.SQLFilter
type SQLRuleAction = original.SQLRuleAction
type Subnet = original.Subnet
type SubscriptionsClient = original.SubscriptionsClient
type TopicsClient = original.TopicsClient
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewArmDisasterRecoveryListResultIterator(page ArmDisasterRecoveryListResultPage) ArmDisasterRecoveryListResultIterator {
	return original.NewArmDisasterRecoveryListResultIterator(page)
}
func NewArmDisasterRecoveryListResultPage(cur ArmDisasterRecoveryListResult, getNextPage func(context.Context, ArmDisasterRecoveryListResult) (ArmDisasterRecoveryListResult, error)) ArmDisasterRecoveryListResultPage {
	return original.NewArmDisasterRecoveryListResultPage(cur, getNextPage)
}
func NewDisasterRecoveryConfigsClient(subscriptionID string) DisasterRecoveryConfigsClient {
	return original.NewDisasterRecoveryConfigsClient(subscriptionID)
}
func NewDisasterRecoveryConfigsClientWithBaseURI(baseURI string, subscriptionID string) DisasterRecoveryConfigsClient {
	return original.NewDisasterRecoveryConfigsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventHubListResultIterator(page EventHubListResultPage) EventHubListResultIterator {
	return original.NewEventHubListResultIterator(page)
}
func NewEventHubListResultPage(cur EventHubListResult, getNextPage func(context.Context, EventHubListResult) (EventHubListResult, error)) EventHubListResultPage {
	return original.NewEventHubListResultPage(cur, getNextPage)
}
func NewEventHubsClient(subscriptionID string) EventHubsClient {
	return original.NewEventHubsClient(subscriptionID)
}
func NewEventHubsClientWithBaseURI(baseURI string, subscriptionID string) EventHubsClient {
	return original.NewEventHubsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMigrationConfigListResultIterator(page MigrationConfigListResultPage) MigrationConfigListResultIterator {
	return original.NewMigrationConfigListResultIterator(page)
}
func NewMigrationConfigListResultPage(cur MigrationConfigListResult, getNextPage func(context.Context, MigrationConfigListResult) (MigrationConfigListResult, error)) MigrationConfigListResultPage {
	return original.NewMigrationConfigListResultPage(cur, getNextPage)
}
func NewMigrationConfigsClient(subscriptionID string) MigrationConfigsClient {
	return original.NewMigrationConfigsClient(subscriptionID)
}
func NewMigrationConfigsClientWithBaseURI(baseURI string, subscriptionID string) MigrationConfigsClient {
	return original.NewMigrationConfigsClientWithBaseURI(baseURI, subscriptionID)
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewNetworkRuleSetListResultIterator(page NetworkRuleSetListResultPage) NetworkRuleSetListResultIterator {
	return original.NewNetworkRuleSetListResultIterator(page)
}
func NewNetworkRuleSetListResultPage(cur NetworkRuleSetListResult, getNextPage func(context.Context, NetworkRuleSetListResult) (NetworkRuleSetListResult, error)) NetworkRuleSetListResultPage {
	return original.NewNetworkRuleSetListResultPage(cur, getNextPage)
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
func NewPremiumMessagingRegionsClient(subscriptionID string) PremiumMessagingRegionsClient {
	return original.NewPremiumMessagingRegionsClient(subscriptionID)
}
func NewPremiumMessagingRegionsClientWithBaseURI(baseURI string, subscriptionID string) PremiumMessagingRegionsClient {
	return original.NewPremiumMessagingRegionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPremiumMessagingRegionsListResultIterator(page PremiumMessagingRegionsListResultPage) PremiumMessagingRegionsListResultIterator {
	return original.NewPremiumMessagingRegionsListResultIterator(page)
}
func NewPremiumMessagingRegionsListResultPage(cur PremiumMessagingRegionsListResult, getNextPage func(context.Context, PremiumMessagingRegionsListResult) (PremiumMessagingRegionsListResult, error)) PremiumMessagingRegionsListResultPage {
	return original.NewPremiumMessagingRegionsListResultPage(cur, getNextPage)
}
func NewQueuesClient(subscriptionID string) QueuesClient {
	return original.NewQueuesClient(subscriptionID)
}
func NewQueuesClientWithBaseURI(baseURI string, subscriptionID string) QueuesClient {
	return original.NewQueuesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegionsClient(subscriptionID string) RegionsClient {
	return original.NewRegionsClient(subscriptionID)
}
func NewRegionsClientWithBaseURI(baseURI string, subscriptionID string) RegionsClient {
	return original.NewRegionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRuleListResultIterator(page RuleListResultPage) RuleListResultIterator {
	return original.NewRuleListResultIterator(page)
}
func NewRuleListResultPage(cur RuleListResult, getNextPage func(context.Context, RuleListResult) (RuleListResult, error)) RuleListResultPage {
	return original.NewRuleListResultPage(cur, getNextPage)
}
func NewRulesClient(subscriptionID string) RulesClient {
	return original.NewRulesClient(subscriptionID)
}
func NewRulesClientWithBaseURI(baseURI string, subscriptionID string) RulesClient {
	return original.NewRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSBAuthorizationRuleListResultIterator(page SBAuthorizationRuleListResultPage) SBAuthorizationRuleListResultIterator {
	return original.NewSBAuthorizationRuleListResultIterator(page)
}
func NewSBAuthorizationRuleListResultPage(cur SBAuthorizationRuleListResult, getNextPage func(context.Context, SBAuthorizationRuleListResult) (SBAuthorizationRuleListResult, error)) SBAuthorizationRuleListResultPage {
	return original.NewSBAuthorizationRuleListResultPage(cur, getNextPage)
}
func NewSBNamespaceListResultIterator(page SBNamespaceListResultPage) SBNamespaceListResultIterator {
	return original.NewSBNamespaceListResultIterator(page)
}
func NewSBNamespaceListResultPage(cur SBNamespaceListResult, getNextPage func(context.Context, SBNamespaceListResult) (SBNamespaceListResult, error)) SBNamespaceListResultPage {
	return original.NewSBNamespaceListResultPage(cur, getNextPage)
}
func NewSBQueueListResultIterator(page SBQueueListResultPage) SBQueueListResultIterator {
	return original.NewSBQueueListResultIterator(page)
}
func NewSBQueueListResultPage(cur SBQueueListResult, getNextPage func(context.Context, SBQueueListResult) (SBQueueListResult, error)) SBQueueListResultPage {
	return original.NewSBQueueListResultPage(cur, getNextPage)
}
func NewSBSubscriptionListResultIterator(page SBSubscriptionListResultPage) SBSubscriptionListResultIterator {
	return original.NewSBSubscriptionListResultIterator(page)
}
func NewSBSubscriptionListResultPage(cur SBSubscriptionListResult, getNextPage func(context.Context, SBSubscriptionListResult) (SBSubscriptionListResult, error)) SBSubscriptionListResultPage {
	return original.NewSBSubscriptionListResultPage(cur, getNextPage)
}
func NewSBTopicListResultIterator(page SBTopicListResultPage) SBTopicListResultIterator {
	return original.NewSBTopicListResultIterator(page)
}
func NewSBTopicListResultPage(cur SBTopicListResult, getNextPage func(context.Context, SBTopicListResult) (SBTopicListResult, error)) SBTopicListResultPage {
	return original.NewSBTopicListResultPage(cur, getNextPage)
}
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClient(subscriptionID)
}
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTopicsClient(subscriptionID string) TopicsClient {
	return original.NewTopicsClient(subscriptionID)
}
func NewTopicsClientWithBaseURI(baseURI string, subscriptionID string) TopicsClient {
	return original.NewTopicsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsValues() []AccessRights {
	return original.PossibleAccessRightsValues()
}
func PossibleDefaultActionValues() []DefaultAction {
	return original.PossibleDefaultActionValues()
}
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return original.PossibleEncodingCaptureDescriptionValues()
}
func PossibleEntityStatusValues() []EntityStatus {
	return original.PossibleEntityStatusValues()
}
func PossibleFilterTypeValues() []FilterType {
	return original.PossibleFilterTypeValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleNameSpaceTypeValues() []NameSpaceType {
	return original.PossibleNameSpaceTypeValues()
}
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return original.PossibleNetworkRuleIPActionValues()
}
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return original.PossibleProvisioningStateDRValues()
}
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return original.PossibleRoleDisasterRecoveryValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleUnavailableReasonValues() []UnavailableReason {
	return original.PossibleUnavailableReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
