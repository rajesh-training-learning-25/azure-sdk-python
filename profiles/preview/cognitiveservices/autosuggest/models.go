//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package autosuggest

import original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/autosuggest"

const (
	DefaultEndpoint = original.DefaultEndpoint
)

type ErrorCode = original.ErrorCode

const (
	InsufficientAuthorization ErrorCode = original.InsufficientAuthorization
	InvalidAuthorization      ErrorCode = original.InvalidAuthorization
	InvalidRequest            ErrorCode = original.InvalidRequest
	None                      ErrorCode = original.None
	RateLimitExceeded         ErrorCode = original.RateLimitExceeded
	ServerError               ErrorCode = original.ServerError
)

type ResponseFormat = original.ResponseFormat

const (
	JSON   ResponseFormat = original.JSON
	JSONLd ResponseFormat = original.JSONLd
)

type SafeSearch = original.SafeSearch

const (
	Moderate SafeSearch = original.Moderate
	Off      SafeSearch = original.Off
	Strict   SafeSearch = original.Strict
)

type ScenarioType = original.ScenarioType

const (
	Custom                  ScenarioType = original.Custom
	PersonalSearchDocuments ScenarioType = original.PersonalSearchDocuments
	PersonalSearchTags      ScenarioType = original.PersonalSearchTags
	SearchHistory           ScenarioType = original.SearchHistory
	StoreApps               ScenarioType = original.StoreApps
	Unknown                 ScenarioType = original.Unknown
	Web                     ScenarioType = original.Web
)

type SearchKind = original.SearchKind

const (
	CustomSearch   SearchKind = original.CustomSearch
	DocumentSearch SearchKind = original.DocumentSearch
	HistorySearch  SearchKind = original.HistorySearch
	LocationSearch SearchKind = original.LocationSearch
	TagSearch      SearchKind = original.TagSearch
	WebSearch      SearchKind = original.WebSearch
)

type Type = original.Type

const (
	TypeSuggestionsSuggestionGroup Type = original.TypeSuggestionsSuggestionGroup
)

type TypeBasicError = original.TypeBasicError

const (
	TypeError TypeBasicError = original.TypeError
)

type TypeBasicQueryContext = original.TypeBasicQueryContext

const (
	TypeQueryContext TypeBasicQueryContext = original.TypeQueryContext
)

type TypeBasicResponseBase = original.TypeBasicResponseBase

const (
	TypeAction              TypeBasicResponseBase = original.TypeAction
	TypeAnswer              TypeBasicResponseBase = original.TypeAnswer
	TypeCreativeWork        TypeBasicResponseBase = original.TypeCreativeWork
	TypeErrorResponse       TypeBasicResponseBase = original.TypeErrorResponse
	TypeIdentifiable        TypeBasicResponseBase = original.TypeIdentifiable
	TypeResponse            TypeBasicResponseBase = original.TypeResponse
	TypeResponseBase        TypeBasicResponseBase = original.TypeResponseBase
	TypeSearchAction        TypeBasicResponseBase = original.TypeSearchAction
	TypeSearchResultsAnswer TypeBasicResponseBase = original.TypeSearchResultsAnswer
	TypeSuggestions         TypeBasicResponseBase = original.TypeSuggestions
	TypeThing               TypeBasicResponseBase = original.TypeThing
)

type Action = original.Action
type Answer = original.Answer
type BaseClient = original.BaseClient
type BasicAction = original.BasicAction
type BasicAnswer = original.BasicAnswer
type BasicCreativeWork = original.BasicCreativeWork
type BasicError = original.BasicError
type BasicIdentifiable = original.BasicIdentifiable
type BasicQueryContext = original.BasicQueryContext
type BasicResponse = original.BasicResponse
type BasicResponseBase = original.BasicResponseBase
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type BasicSuggestionsSuggestionGroup = original.BasicSuggestionsSuggestionGroup
type BasicThing = original.BasicThing
type CreativeWork = original.CreativeWork
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type Identifiable = original.Identifiable
type QueryContext = original.QueryContext
type Response = original.Response
type ResponseBase = original.ResponseBase
type SearchAction = original.SearchAction
type SearchResultsAnswer = original.SearchResultsAnswer
type Suggestions = original.Suggestions
type SuggestionsSuggestionGroup = original.SuggestionsSuggestionGroup
type Thing = original.Thing

func New() BaseClient {
	return original.New()
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleResponseFormatValues() []ResponseFormat {
	return original.PossibleResponseFormatValues()
}
func PossibleSafeSearchValues() []SafeSearch {
	return original.PossibleSafeSearchValues()
}
func PossibleScenarioTypeValues() []ScenarioType {
	return original.PossibleScenarioTypeValues()
}
func PossibleSearchKindValues() []SearchKind {
	return original.PossibleSearchKindValues()
}
func PossibleTypeBasicErrorValues() []TypeBasicError {
	return original.PossibleTypeBasicErrorValues()
}
func PossibleTypeBasicQueryContextValues() []TypeBasicQueryContext {
	return original.PossibleTypeBasicQueryContextValues()
}
func PossibleTypeBasicResponseBaseValues() []TypeBasicResponseBase {
	return original.PossibleTypeBasicResponseBaseValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
