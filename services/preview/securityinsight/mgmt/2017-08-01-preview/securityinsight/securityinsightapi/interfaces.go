package securityinsightapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/securityinsight/mgmt/2017-08-01-preview/securityinsight"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	ListDataConnectorRequirements(ctx context.Context, resourceGroupName string, workspaceName string, operationalInsightsResourceProvider string, dataConnectorsCheckRequirements securityinsight.DataConnectorsCheckRequirements) (result securityinsight.DataConnectorRequirementsState, err error)
}

var _ BaseClientAPI = (*securityinsight.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result securityinsight.OperationsListPage, err error)
	ListComplete(ctx context.Context) (result securityinsight.OperationsListIterator, err error)
}

var _ OperationsClientAPI = (*securityinsight.OperationsClient)(nil)

// AlertRulesClientAPI contains the set of methods on the AlertRulesClient type.
type AlertRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, alertRule securityinsight.BasicAlertRule) (result securityinsight.AlertRuleModel, err error)
	CreateOrUpdateAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string, action securityinsight.ActionRequest) (result securityinsight.ActionResponse, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result autorest.Response, err error)
	DeleteAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result securityinsight.AlertRuleModel, err error)
	GetAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string) (result securityinsight.ActionResponse, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.AlertRulesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.AlertRulesListIterator, err error)
}

var _ AlertRulesClientAPI = (*securityinsight.AlertRulesClient)(nil)

// ActionsClientAPI contains the set of methods on the ActionsClient type.
type ActionsClientAPI interface {
	ListByAlertRule(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result securityinsight.ActionsListPage, err error)
	ListByAlertRuleComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result securityinsight.ActionsListIterator, err error)
}

var _ ActionsClientAPI = (*securityinsight.ActionsClient)(nil)

// AlertRuleTemplatesClientAPI contains the set of methods on the AlertRuleTemplatesClient type.
type AlertRuleTemplatesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, alertRuleTemplateID string) (result securityinsight.AlertRuleTemplateModel, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.AlertRuleTemplatesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.AlertRuleTemplatesListIterator, err error)
}

var _ AlertRuleTemplatesClientAPI = (*securityinsight.AlertRuleTemplatesClient)(nil)

// CasesClientAPI contains the set of methods on the CasesClient type.
type CasesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseParameter securityinsight.Case) (result securityinsight.Case, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string) (result securityinsight.Case, err error)
	GetComment(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseCommentID string) (result securityinsight.CaseComment, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.CaseListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.CaseListIterator, err error)
}

var _ CasesClientAPI = (*securityinsight.CasesClient)(nil)

// CommentsClientAPI contains the set of methods on the CommentsClient type.
type CommentsClientAPI interface {
	ListByCase(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.CaseCommentListPage, err error)
	ListByCaseComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.CaseCommentListIterator, err error)
}

var _ CommentsClientAPI = (*securityinsight.CommentsClient)(nil)

// CaseCommentsClientAPI contains the set of methods on the CaseCommentsClient type.
type CaseCommentsClientAPI interface {
	CreateComment(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseCommentID string, caseComment securityinsight.CaseComment) (result securityinsight.CaseComment, err error)
}

var _ CaseCommentsClientAPI = (*securityinsight.CaseCommentsClient)(nil)

// BookmarksClientAPI contains the set of methods on the BookmarksClient type.
type BookmarksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, bookmark securityinsight.Bookmark) (result securityinsight.Bookmark, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string) (result securityinsight.Bookmark, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.BookmarkListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.BookmarkListIterator, err error)
}

var _ BookmarksClientAPI = (*securityinsight.BookmarksClient)(nil)

// CaseRelationsClientAPI contains the set of methods on the CaseRelationsClient type.
type CaseRelationsClientAPI interface {
	CreateOrUpdateRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string, relationInputModel securityinsight.RelationsModelInput) (result securityinsight.CaseRelation, err error)
	DeleteRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string) (result autorest.Response, err error)
	GetRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, relationName string) (result securityinsight.CaseRelation, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.CaseRelationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.CaseRelationListIterator, err error)
}

var _ CaseRelationsClientAPI = (*securityinsight.CaseRelationsClient)(nil)

// BookmarkRelationsClientAPI contains the set of methods on the BookmarkRelationsClient type.
type BookmarkRelationsClientAPI interface {
	CreateOrUpdateRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string, relationInputModel securityinsight.RelationsModelInput) (result securityinsight.BookmarkRelation, err error)
	DeleteRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string) (result autorest.Response, err error)
	GetRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string) (result securityinsight.BookmarkRelation, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.BookmarkRelationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.BookmarkRelationListIterator, err error)
}

var _ BookmarkRelationsClientAPI = (*securityinsight.BookmarkRelationsClient)(nil)

// DataConnectorsClientAPI contains the set of methods on the DataConnectorsClient type.
type DataConnectorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, dataConnectorID string, dataConnector securityinsight.BasicDataConnector) (result securityinsight.DataConnectorModel, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, dataConnectorID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, dataConnectorID string) (result securityinsight.DataConnectorModel, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.DataConnectorListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.DataConnectorListIterator, err error)
}

var _ DataConnectorsClientAPI = (*securityinsight.DataConnectorsClient)(nil)

// EntitiesClientAPI contains the set of methods on the EntitiesClient type.
type EntitiesClientAPI interface {
	Expand(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityID string, parameters securityinsight.EntityExpandParameters) (result securityinsight.EntityExpandResponse, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityID string) (result securityinsight.EntityModel, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.EntityListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.EntityListIterator, err error)
}

var _ EntitiesClientAPI = (*securityinsight.EntitiesClient)(nil)

// EntitiesRelationsClientAPI contains the set of methods on the EntitiesRelationsClient type.
type EntitiesRelationsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListIterator, err error)
}

var _ EntitiesRelationsClientAPI = (*securityinsight.EntitiesRelationsClient)(nil)

// EntityRelationsClientAPI contains the set of methods on the EntityRelationsClient type.
type EntityRelationsClientAPI interface {
	GetRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityID string, relationName string) (result securityinsight.Relation, err error)
}

var _ EntityRelationsClientAPI = (*securityinsight.EntityRelationsClient)(nil)

// OfficeConsentsClientAPI contains the set of methods on the OfficeConsentsClient type.
type OfficeConsentsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, consentID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, consentID string) (result securityinsight.OfficeConsent, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.OfficeConsentListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.OfficeConsentListIterator, err error)
}

var _ OfficeConsentsClientAPI = (*securityinsight.OfficeConsentsClient)(nil)

// ProductSettingsClientAPI contains the set of methods on the ProductSettingsClient type.
type ProductSettingsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string) (result securityinsight.SettingsModel, err error)
	Update(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string, settings securityinsight.BasicSettings) (result securityinsight.SettingsModel, err error)
}

var _ ProductSettingsClientAPI = (*securityinsight.ProductSettingsClient)(nil)

// CasesAggregationsClientAPI contains the set of methods on the CasesAggregationsClient type.
type CasesAggregationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, aggregationsName string) (result securityinsight.AggregationsModel, err error)
}

var _ CasesAggregationsClientAPI = (*securityinsight.CasesAggregationsClient)(nil)

// EntityQueriesClientAPI contains the set of methods on the EntityQueriesClient type.
type EntityQueriesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityQueryID string) (result securityinsight.EntityQuery, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.EntityQueryListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.EntityQueryListIterator, err error)
}

var _ EntityQueriesClientAPI = (*securityinsight.EntityQueriesClient)(nil)

// IncidentsClientAPI contains the set of methods on the IncidentsClient type.
type IncidentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, incident securityinsight.Incident) (result securityinsight.Incident, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string) (result securityinsight.Incident, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentListIterator, err error)
}

var _ IncidentsClientAPI = (*securityinsight.IncidentsClient)(nil)

// IncidentCommentsClientAPI contains the set of methods on the IncidentCommentsClient type.
type IncidentCommentsClientAPI interface {
	CreateComment(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, incidentCommentID string, incidentComment securityinsight.IncidentComment) (result securityinsight.IncidentComment, err error)
	GetComment(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, incidentCommentID string) (result securityinsight.IncidentComment, err error)
	ListByIncident(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentCommentListPage, err error)
	ListByIncidentComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentCommentListIterator, err error)
}

var _ IncidentCommentsClientAPI = (*securityinsight.IncidentCommentsClient)(nil)

// IncidentRelationsClientAPI contains the set of methods on the IncidentRelationsClient type.
type IncidentRelationsClientAPI interface {
	CreateOrUpdateRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, relationName string, relation securityinsight.Relation) (result securityinsight.Relation, err error)
	DeleteRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, relationName string) (result autorest.Response, err error)
	GetRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, relationName string) (result securityinsight.Relation, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.RelationListIterator, err error)
}

var _ IncidentRelationsClientAPI = (*securityinsight.IncidentRelationsClient)(nil)
