//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package insights

import (
	"context"

	original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/monitor/mgmt/2020-10-01/insights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionGroup = original.ActionGroup
type ActionList = original.ActionList
type ActivityLogAlertResource = original.ActivityLogAlertResource
type ActivityLogAlertsClient = original.ActivityLogAlertsClient
type AlertRuleAllOfCondition = original.AlertRuleAllOfCondition
type AlertRuleAnyOfOrLeafCondition = original.AlertRuleAnyOfOrLeafCondition
type AlertRuleLeafCondition = original.AlertRuleLeafCondition
type AlertRuleList = original.AlertRuleList
type AlertRuleListIterator = original.AlertRuleListIterator
type AlertRuleListPage = original.AlertRuleListPage
type AlertRulePatchObject = original.AlertRulePatchObject
type AlertRulePatchProperties = original.AlertRulePatchProperties
type AlertRuleProperties = original.AlertRuleProperties
type AzureResource = original.AzureResource
type BaseClient = original.BaseClient
type ErrorResponse = original.ErrorResponse

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActivityLogAlertsClient(subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClient(subscriptionID)
}
func NewActivityLogAlertsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRuleListIterator(page AlertRuleListPage) AlertRuleListIterator {
	return original.NewAlertRuleListIterator(page)
}
func NewAlertRuleListPage(cur AlertRuleList, getNextPage func(context.Context, AlertRuleList) (AlertRuleList, error)) AlertRuleListPage {
	return original.NewAlertRuleListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
