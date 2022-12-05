//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights

// ActionsClientCreateOrUpdateResponse contains the response from method ActionsClient.CreateOrUpdate.
type ActionsClientCreateOrUpdateResponse struct {
	ActionResponse
}

// ActionsClientDeleteResponse contains the response from method ActionsClient.Delete.
type ActionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActionsClientGetResponse contains the response from method ActionsClient.Get.
type ActionsClientGetResponse struct {
	ActionResponse
}

// ActionsClientListByAlertRuleResponse contains the response from method ActionsClient.ListByAlertRule.
type ActionsClientListByAlertRuleResponse struct {
	ActionsList
}

// AlertRuleTemplatesClientGetResponse contains the response from method AlertRuleTemplatesClient.Get.
type AlertRuleTemplatesClientGetResponse struct {
	AlertRuleTemplateClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AlertRuleTemplatesClientGetResponse.
func (a *AlertRuleTemplatesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAlertRuleTemplateClassification(data)
	if err != nil {
		return err
	}
	a.AlertRuleTemplateClassification = res
	return nil
}

// AlertRuleTemplatesClientListResponse contains the response from method AlertRuleTemplatesClient.List.
type AlertRuleTemplatesClientListResponse struct {
	AlertRuleTemplatesList
}

// AlertRulesClientCreateOrUpdateResponse contains the response from method AlertRulesClient.CreateOrUpdate.
type AlertRulesClientCreateOrUpdateResponse struct {
	AlertRuleClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AlertRulesClientCreateOrUpdateResponse.
func (a *AlertRulesClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAlertRuleClassification(data)
	if err != nil {
		return err
	}
	a.AlertRuleClassification = res
	return nil
}

// AlertRulesClientDeleteResponse contains the response from method AlertRulesClient.Delete.
type AlertRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// AlertRulesClientGetResponse contains the response from method AlertRulesClient.Get.
type AlertRulesClientGetResponse struct {
	AlertRuleClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AlertRulesClientGetResponse.
func (a *AlertRulesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAlertRuleClassification(data)
	if err != nil {
		return err
	}
	a.AlertRuleClassification = res
	return nil
}

// AlertRulesClientListResponse contains the response from method AlertRulesClient.List.
type AlertRulesClientListResponse struct {
	AlertRulesList
}

// AutomationRulesClientCreateOrUpdateResponse contains the response from method AutomationRulesClient.CreateOrUpdate.
type AutomationRulesClientCreateOrUpdateResponse struct {
	AutomationRule
}

// AutomationRulesClientDeleteResponse contains the response from method AutomationRulesClient.Delete.
type AutomationRulesClientDeleteResponse struct {
	// Anything
	Interface interface{}
}

// AutomationRulesClientGetResponse contains the response from method AutomationRulesClient.Get.
type AutomationRulesClientGetResponse struct {
	AutomationRule
}

// AutomationRulesClientListResponse contains the response from method AutomationRulesClient.List.
type AutomationRulesClientListResponse struct {
	AutomationRulesList
}

// BookmarkClientExpandResponse contains the response from method BookmarkClient.Expand.
type BookmarkClientExpandResponse struct {
	BookmarkExpandResponse
}

// BookmarkRelationsClientCreateOrUpdateResponse contains the response from method BookmarkRelationsClient.CreateOrUpdate.
type BookmarkRelationsClientCreateOrUpdateResponse struct {
	Relation
}

// BookmarkRelationsClientDeleteResponse contains the response from method BookmarkRelationsClient.Delete.
type BookmarkRelationsClientDeleteResponse struct {
	// placeholder for future response values
}

// BookmarkRelationsClientGetResponse contains the response from method BookmarkRelationsClient.Get.
type BookmarkRelationsClientGetResponse struct {
	Relation
}

// BookmarkRelationsClientListResponse contains the response from method BookmarkRelationsClient.List.
type BookmarkRelationsClientListResponse struct {
	RelationList
}

// BookmarksClientCreateOrUpdateResponse contains the response from method BookmarksClient.CreateOrUpdate.
type BookmarksClientCreateOrUpdateResponse struct {
	Bookmark
}

// BookmarksClientDeleteResponse contains the response from method BookmarksClient.Delete.
type BookmarksClientDeleteResponse struct {
	// placeholder for future response values
}

// BookmarksClientGetResponse contains the response from method BookmarksClient.Get.
type BookmarksClientGetResponse struct {
	Bookmark
}

// BookmarksClientListResponse contains the response from method BookmarksClient.List.
type BookmarksClientListResponse struct {
	BookmarkList
}

// DataConnectorsCheckRequirementsClientPostResponse contains the response from method DataConnectorsCheckRequirementsClient.Post.
type DataConnectorsCheckRequirementsClientPostResponse struct {
	DataConnectorRequirementsState
}

// DataConnectorsClientConnectResponse contains the response from method DataConnectorsClient.Connect.
type DataConnectorsClientConnectResponse struct {
	// placeholder for future response values
}

// DataConnectorsClientCreateOrUpdateResponse contains the response from method DataConnectorsClient.CreateOrUpdate.
type DataConnectorsClientCreateOrUpdateResponse struct {
	DataConnectorClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataConnectorsClientCreateOrUpdateResponse.
func (d *DataConnectorsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataConnectorClassification(data)
	if err != nil {
		return err
	}
	d.DataConnectorClassification = res
	return nil
}

// DataConnectorsClientDeleteResponse contains the response from method DataConnectorsClient.Delete.
type DataConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataConnectorsClientDisconnectResponse contains the response from method DataConnectorsClient.Disconnect.
type DataConnectorsClientDisconnectResponse struct {
	// placeholder for future response values
}

// DataConnectorsClientGetResponse contains the response from method DataConnectorsClient.Get.
type DataConnectorsClientGetResponse struct {
	DataConnectorClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataConnectorsClientGetResponse.
func (d *DataConnectorsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataConnectorClassification(data)
	if err != nil {
		return err
	}
	d.DataConnectorClassification = res
	return nil
}

// DataConnectorsClientListResponse contains the response from method DataConnectorsClient.List.
type DataConnectorsClientListResponse struct {
	DataConnectorList
}

// DomainWhoisClientGetResponse contains the response from method DomainWhoisClient.Get.
type DomainWhoisClientGetResponse struct {
	EnrichmentDomainWhois
}

// EntitiesClientExpandResponse contains the response from method EntitiesClient.Expand.
type EntitiesClientExpandResponse struct {
	EntityExpandResponse
}

// EntitiesClientGetInsightsResponse contains the response from method EntitiesClient.GetInsights.
type EntitiesClientGetInsightsResponse struct {
	EntityGetInsightsResponse
}

// EntitiesClientGetResponse contains the response from method EntitiesClient.Get.
type EntitiesClientGetResponse struct {
	EntityClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EntitiesClientGetResponse.
func (e *EntitiesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEntityClassification(data)
	if err != nil {
		return err
	}
	e.EntityClassification = res
	return nil
}

// EntitiesClientListResponse contains the response from method EntitiesClient.List.
type EntitiesClientListResponse struct {
	EntityList
}

// EntitiesClientQueriesResponse contains the response from method EntitiesClient.Queries.
type EntitiesClientQueriesResponse struct {
	GetQueriesResponse
}

// EntitiesGetTimelineClientListResponse contains the response from method EntitiesGetTimelineClient.List.
type EntitiesGetTimelineClientListResponse struct {
	EntityTimelineResponse
}

// EntitiesRelationsClientListResponse contains the response from method EntitiesRelationsClient.List.
type EntitiesRelationsClientListResponse struct {
	RelationList
}

// EntityQueriesClientCreateOrUpdateResponse contains the response from method EntityQueriesClient.CreateOrUpdate.
type EntityQueriesClientCreateOrUpdateResponse struct {
	EntityQueryClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EntityQueriesClientCreateOrUpdateResponse.
func (e *EntityQueriesClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEntityQueryClassification(data)
	if err != nil {
		return err
	}
	e.EntityQueryClassification = res
	return nil
}

// EntityQueriesClientDeleteResponse contains the response from method EntityQueriesClient.Delete.
type EntityQueriesClientDeleteResponse struct {
	// placeholder for future response values
}

// EntityQueriesClientGetResponse contains the response from method EntityQueriesClient.Get.
type EntityQueriesClientGetResponse struct {
	EntityQueryClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EntityQueriesClientGetResponse.
func (e *EntityQueriesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEntityQueryClassification(data)
	if err != nil {
		return err
	}
	e.EntityQueryClassification = res
	return nil
}

// EntityQueriesClientListResponse contains the response from method EntityQueriesClient.List.
type EntityQueriesClientListResponse struct {
	EntityQueryList
}

// EntityQueryTemplatesClientGetResponse contains the response from method EntityQueryTemplatesClient.Get.
type EntityQueryTemplatesClientGetResponse struct {
	EntityQueryTemplateClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EntityQueryTemplatesClientGetResponse.
func (e *EntityQueryTemplatesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEntityQueryTemplateClassification(data)
	if err != nil {
		return err
	}
	e.EntityQueryTemplateClassification = res
	return nil
}

// EntityQueryTemplatesClientListResponse contains the response from method EntityQueryTemplatesClient.List.
type EntityQueryTemplatesClientListResponse struct {
	EntityQueryTemplateList
}

// EntityRelationsClientGetRelationResponse contains the response from method EntityRelationsClient.GetRelation.
type EntityRelationsClientGetRelationResponse struct {
	Relation
}

// FileImportsClientCreateResponse contains the response from method FileImportsClient.Create.
type FileImportsClientCreateResponse struct {
	FileImport
}

// FileImportsClientDeleteResponse contains the response from method FileImportsClient.Delete.
type FileImportsClientDeleteResponse struct {
	FileImport
}

// FileImportsClientGetResponse contains the response from method FileImportsClient.Get.
type FileImportsClientGetResponse struct {
	FileImport
}

// FileImportsClientListResponse contains the response from method FileImportsClient.List.
type FileImportsClientListResponse struct {
	FileImportList
}

// IPGeodataClientGetResponse contains the response from method IPGeodataClient.Get.
type IPGeodataClientGetResponse struct {
	EnrichmentIPGeodata
}

// IncidentCommentsClientCreateOrUpdateResponse contains the response from method IncidentCommentsClient.CreateOrUpdate.
type IncidentCommentsClientCreateOrUpdateResponse struct {
	IncidentComment
}

// IncidentCommentsClientDeleteResponse contains the response from method IncidentCommentsClient.Delete.
type IncidentCommentsClientDeleteResponse struct {
	// placeholder for future response values
}

// IncidentCommentsClientGetResponse contains the response from method IncidentCommentsClient.Get.
type IncidentCommentsClientGetResponse struct {
	IncidentComment
}

// IncidentCommentsClientListResponse contains the response from method IncidentCommentsClient.List.
type IncidentCommentsClientListResponse struct {
	IncidentCommentList
}

// IncidentRelationsClientCreateOrUpdateResponse contains the response from method IncidentRelationsClient.CreateOrUpdate.
type IncidentRelationsClientCreateOrUpdateResponse struct {
	Relation
}

// IncidentRelationsClientDeleteResponse contains the response from method IncidentRelationsClient.Delete.
type IncidentRelationsClientDeleteResponse struct {
	// placeholder for future response values
}

// IncidentRelationsClientGetResponse contains the response from method IncidentRelationsClient.Get.
type IncidentRelationsClientGetResponse struct {
	Relation
}

// IncidentRelationsClientListResponse contains the response from method IncidentRelationsClient.List.
type IncidentRelationsClientListResponse struct {
	RelationList
}

// IncidentsClientCreateOrUpdateResponse contains the response from method IncidentsClient.CreateOrUpdate.
type IncidentsClientCreateOrUpdateResponse struct {
	Incident
}

// IncidentsClientCreateTeamResponse contains the response from method IncidentsClient.CreateTeam.
type IncidentsClientCreateTeamResponse struct {
	TeamInformation
}

// IncidentsClientDeleteResponse contains the response from method IncidentsClient.Delete.
type IncidentsClientDeleteResponse struct {
	// placeholder for future response values
}

// IncidentsClientGetResponse contains the response from method IncidentsClient.Get.
type IncidentsClientGetResponse struct {
	Incident
}

// IncidentsClientListAlertsResponse contains the response from method IncidentsClient.ListAlerts.
type IncidentsClientListAlertsResponse struct {
	IncidentAlertList
}

// IncidentsClientListBookmarksResponse contains the response from method IncidentsClient.ListBookmarks.
type IncidentsClientListBookmarksResponse struct {
	IncidentBookmarkList
}

// IncidentsClientListEntitiesResponse contains the response from method IncidentsClient.ListEntities.
type IncidentsClientListEntitiesResponse struct {
	IncidentEntitiesResponse
}

// IncidentsClientListResponse contains the response from method IncidentsClient.List.
type IncidentsClientListResponse struct {
	IncidentList
}

// IncidentsClientRunPlaybookResponse contains the response from method IncidentsClient.RunPlaybook.
type IncidentsClientRunPlaybookResponse struct {
	// Anything
	Interface interface{}
}

// MetadataClientCreateResponse contains the response from method MetadataClient.Create.
type MetadataClientCreateResponse struct {
	MetadataModel
}

// MetadataClientDeleteResponse contains the response from method MetadataClient.Delete.
type MetadataClientDeleteResponse struct {
	// placeholder for future response values
}

// MetadataClientGetResponse contains the response from method MetadataClient.Get.
type MetadataClientGetResponse struct {
	MetadataModel
}

// MetadataClientListResponse contains the response from method MetadataClient.List.
type MetadataClientListResponse struct {
	MetadataList
}

// MetadataClientUpdateResponse contains the response from method MetadataClient.Update.
type MetadataClientUpdateResponse struct {
	MetadataModel
}

// OfficeConsentsClientDeleteResponse contains the response from method OfficeConsentsClient.Delete.
type OfficeConsentsClientDeleteResponse struct {
	// placeholder for future response values
}

// OfficeConsentsClientGetResponse contains the response from method OfficeConsentsClient.Get.
type OfficeConsentsClientGetResponse struct {
	OfficeConsent
}

// OfficeConsentsClientListResponse contains the response from method OfficeConsentsClient.List.
type OfficeConsentsClientListResponse struct {
	OfficeConsentList
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsList
}

// ProductSettingsClientDeleteResponse contains the response from method ProductSettingsClient.Delete.
type ProductSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProductSettingsClientGetResponse contains the response from method ProductSettingsClient.Get.
type ProductSettingsClientGetResponse struct {
	SettingsClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProductSettingsClientGetResponse.
func (p *ProductSettingsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSettingsClassification(data)
	if err != nil {
		return err
	}
	p.SettingsClassification = res
	return nil
}

// ProductSettingsClientListResponse contains the response from method ProductSettingsClient.List.
type ProductSettingsClientListResponse struct {
	SettingList
}

// ProductSettingsClientUpdateResponse contains the response from method ProductSettingsClient.Update.
type ProductSettingsClientUpdateResponse struct {
	SettingsClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProductSettingsClientUpdateResponse.
func (p *ProductSettingsClientUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSettingsClassification(data)
	if err != nil {
		return err
	}
	p.SettingsClassification = res
	return nil
}

// SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse contains the response from method SecurityMLAnalyticsSettingsClient.CreateOrUpdate.
type SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse struct {
	SecurityMLAnalyticsSettingClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse.
func (s *SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSecurityMLAnalyticsSettingClassification(data)
	if err != nil {
		return err
	}
	s.SecurityMLAnalyticsSettingClassification = res
	return nil
}

// SecurityMLAnalyticsSettingsClientDeleteResponse contains the response from method SecurityMLAnalyticsSettingsClient.Delete.
type SecurityMLAnalyticsSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// SecurityMLAnalyticsSettingsClientGetResponse contains the response from method SecurityMLAnalyticsSettingsClient.Get.
type SecurityMLAnalyticsSettingsClientGetResponse struct {
	SecurityMLAnalyticsSettingClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecurityMLAnalyticsSettingsClientGetResponse.
func (s *SecurityMLAnalyticsSettingsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSecurityMLAnalyticsSettingClassification(data)
	if err != nil {
		return err
	}
	s.SecurityMLAnalyticsSettingClassification = res
	return nil
}

// SecurityMLAnalyticsSettingsClientListResponse contains the response from method SecurityMLAnalyticsSettingsClient.List.
type SecurityMLAnalyticsSettingsClientListResponse struct {
	SecurityMLAnalyticsSettingsList
}

// SentinelOnboardingStatesClientCreateResponse contains the response from method SentinelOnboardingStatesClient.Create.
type SentinelOnboardingStatesClientCreateResponse struct {
	SentinelOnboardingState
}

// SentinelOnboardingStatesClientDeleteResponse contains the response from method SentinelOnboardingStatesClient.Delete.
type SentinelOnboardingStatesClientDeleteResponse struct {
	// placeholder for future response values
}

// SentinelOnboardingStatesClientGetResponse contains the response from method SentinelOnboardingStatesClient.Get.
type SentinelOnboardingStatesClientGetResponse struct {
	SentinelOnboardingState
}

// SentinelOnboardingStatesClientListResponse contains the response from method SentinelOnboardingStatesClient.List.
type SentinelOnboardingStatesClientListResponse struct {
	SentinelOnboardingStatesList
}

// SourceControlClientListRepositoriesResponse contains the response from method SourceControlClient.ListRepositories.
type SourceControlClientListRepositoriesResponse struct {
	RepoList
}

// SourceControlsClientCreateResponse contains the response from method SourceControlsClient.Create.
type SourceControlsClientCreateResponse struct {
	SourceControl
}

// SourceControlsClientDeleteResponse contains the response from method SourceControlsClient.Delete.
type SourceControlsClientDeleteResponse struct {
	// placeholder for future response values
}

// SourceControlsClientGetResponse contains the response from method SourceControlsClient.Get.
type SourceControlsClientGetResponse struct {
	SourceControl
}

// SourceControlsClientListResponse contains the response from method SourceControlsClient.List.
type SourceControlsClientListResponse struct {
	SourceControlList
}

// ThreatIntelligenceIndicatorClientAppendTagsResponse contains the response from method ThreatIntelligenceIndicatorClient.AppendTags.
type ThreatIntelligenceIndicatorClientAppendTagsResponse struct {
	// placeholder for future response values
}

// ThreatIntelligenceIndicatorClientCreateIndicatorResponse contains the response from method ThreatIntelligenceIndicatorClient.CreateIndicator.
type ThreatIntelligenceIndicatorClientCreateIndicatorResponse struct {
	ThreatIntelligenceInformationClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreatIntelligenceIndicatorClientCreateIndicatorResponse.
func (t *ThreatIntelligenceIndicatorClientCreateIndicatorResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalThreatIntelligenceInformationClassification(data)
	if err != nil {
		return err
	}
	t.ThreatIntelligenceInformationClassification = res
	return nil
}

// ThreatIntelligenceIndicatorClientCreateResponse contains the response from method ThreatIntelligenceIndicatorClient.Create.
type ThreatIntelligenceIndicatorClientCreateResponse struct {
	ThreatIntelligenceInformationClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreatIntelligenceIndicatorClientCreateResponse.
func (t *ThreatIntelligenceIndicatorClientCreateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalThreatIntelligenceInformationClassification(data)
	if err != nil {
		return err
	}
	t.ThreatIntelligenceInformationClassification = res
	return nil
}

// ThreatIntelligenceIndicatorClientDeleteResponse contains the response from method ThreatIntelligenceIndicatorClient.Delete.
type ThreatIntelligenceIndicatorClientDeleteResponse struct {
	// placeholder for future response values
}

// ThreatIntelligenceIndicatorClientGetResponse contains the response from method ThreatIntelligenceIndicatorClient.Get.
type ThreatIntelligenceIndicatorClientGetResponse struct {
	ThreatIntelligenceInformationClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreatIntelligenceIndicatorClientGetResponse.
func (t *ThreatIntelligenceIndicatorClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalThreatIntelligenceInformationClassification(data)
	if err != nil {
		return err
	}
	t.ThreatIntelligenceInformationClassification = res
	return nil
}

// ThreatIntelligenceIndicatorClientQueryIndicatorsResponse contains the response from method ThreatIntelligenceIndicatorClient.QueryIndicators.
type ThreatIntelligenceIndicatorClientQueryIndicatorsResponse struct {
	ThreatIntelligenceInformationList
}

// ThreatIntelligenceIndicatorClientReplaceTagsResponse contains the response from method ThreatIntelligenceIndicatorClient.ReplaceTags.
type ThreatIntelligenceIndicatorClientReplaceTagsResponse struct {
	ThreatIntelligenceInformationClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreatIntelligenceIndicatorClientReplaceTagsResponse.
func (t *ThreatIntelligenceIndicatorClientReplaceTagsResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalThreatIntelligenceInformationClassification(data)
	if err != nil {
		return err
	}
	t.ThreatIntelligenceInformationClassification = res
	return nil
}

// ThreatIntelligenceIndicatorMetricsClientListResponse contains the response from method ThreatIntelligenceIndicatorMetricsClient.List.
type ThreatIntelligenceIndicatorMetricsClientListResponse struct {
	ThreatIntelligenceMetricsList
}

// ThreatIntelligenceIndicatorsClientListResponse contains the response from method ThreatIntelligenceIndicatorsClient.List.
type ThreatIntelligenceIndicatorsClientListResponse struct {
	ThreatIntelligenceInformationList
}

// WatchlistItemsClientCreateOrUpdateResponse contains the response from method WatchlistItemsClient.CreateOrUpdate.
type WatchlistItemsClientCreateOrUpdateResponse struct {
	WatchlistItem
}

// WatchlistItemsClientDeleteResponse contains the response from method WatchlistItemsClient.Delete.
type WatchlistItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// WatchlistItemsClientGetResponse contains the response from method WatchlistItemsClient.Get.
type WatchlistItemsClientGetResponse struct {
	WatchlistItem
}

// WatchlistItemsClientListResponse contains the response from method WatchlistItemsClient.List.
type WatchlistItemsClientListResponse struct {
	WatchlistItemList
}

// WatchlistsClientCreateOrUpdateResponse contains the response from method WatchlistsClient.CreateOrUpdate.
type WatchlistsClientCreateOrUpdateResponse struct {
	Watchlist
	// AzureAsyncOperation contains the information returned from the Azure-AsyncOperation header response.
	AzureAsyncOperation *string
}

// WatchlistsClientDeleteResponse contains the response from method WatchlistsClient.Delete.
type WatchlistsClientDeleteResponse struct {
	// AzureAsyncOperation contains the information returned from the Azure-AsyncOperation header response.
	AzureAsyncOperation *string
}

// WatchlistsClientGetResponse contains the response from method WatchlistsClient.Get.
type WatchlistsClientGetResponse struct {
	Watchlist
}

// WatchlistsClientListResponse contains the response from method WatchlistsClient.List.
type WatchlistsClientListResponse struct {
	WatchlistList
}