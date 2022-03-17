//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package localsearch

import original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/localsearch"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type EntityScenario = original.EntityScenario

const (
	DisambiguationItem EntityScenario = original.DisambiguationItem
	DominantEntity     EntityScenario = original.DominantEntity
	ListItem           EntityScenario = original.ListItem
)

type EntityType = original.EntityType

const (
	EntityTypeHotel         EntityType = original.EntityTypeHotel
	EntityTypeLocalBusiness EntityType = original.EntityTypeLocalBusiness
	EntityTypePlace         EntityType = original.EntityTypePlace
	EntityTypeRestaurant    EntityType = original.EntityTypeRestaurant
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

type ErrorSubCode = original.ErrorSubCode

const (
	AuthorizationDisabled   ErrorSubCode = original.AuthorizationDisabled
	AuthorizationExpired    ErrorSubCode = original.AuthorizationExpired
	AuthorizationMissing    ErrorSubCode = original.AuthorizationMissing
	AuthorizationRedundancy ErrorSubCode = original.AuthorizationRedundancy
	Blocked                 ErrorSubCode = original.Blocked
	HTTPNotAllowed          ErrorSubCode = original.HTTPNotAllowed
	NotImplemented          ErrorSubCode = original.NotImplemented
	ParameterInvalidValue   ErrorSubCode = original.ParameterInvalidValue
	ParameterMissing        ErrorSubCode = original.ParameterMissing
	ResourceError           ErrorSubCode = original.ResourceError
	UnexpectedError         ErrorSubCode = original.UnexpectedError
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

type Type = original.Type

const (
	TypeQueryContext Type = original.TypeQueryContext
)

type TypeBasicEntitiesEntityPresentationInfo = original.TypeBasicEntitiesEntityPresentationInfo

const (
	TypeEntitiesEntityPresentationInfo TypeBasicEntitiesEntityPresentationInfo = original.TypeEntitiesEntityPresentationInfo
)

type TypeBasicError = original.TypeBasicError

const (
	TypeError TypeBasicError = original.TypeError
)

type TypeBasicGeoCoordinates = original.TypeBasicGeoCoordinates

const (
	TypeGeoCoordinates TypeBasicGeoCoordinates = original.TypeGeoCoordinates
)

type TypeBasicResponseBase = original.TypeBasicResponseBase

const (
	TypeAction              TypeBasicResponseBase = original.TypeAction
	TypeAnswer              TypeBasicResponseBase = original.TypeAnswer
	TypeCreativeWork        TypeBasicResponseBase = original.TypeCreativeWork
	TypeErrorResponse       TypeBasicResponseBase = original.TypeErrorResponse
	TypeIdentifiable        TypeBasicResponseBase = original.TypeIdentifiable
	TypeIntangible          TypeBasicResponseBase = original.TypeIntangible
	TypePlace               TypeBasicResponseBase = original.TypePlace
	TypePlaces              TypeBasicResponseBase = original.TypePlaces
	TypePostalAddress       TypeBasicResponseBase = original.TypePostalAddress
	TypeResponse            TypeBasicResponseBase = original.TypeResponse
	TypeResponseBase        TypeBasicResponseBase = original.TypeResponseBase
	TypeSearchAction        TypeBasicResponseBase = original.TypeSearchAction
	TypeSearchResponse      TypeBasicResponseBase = original.TypeSearchResponse
	TypeSearchResultsAnswer TypeBasicResponseBase = original.TypeSearchResultsAnswer
	TypeStructuredValue     TypeBasicResponseBase = original.TypeStructuredValue
	TypeThing               TypeBasicResponseBase = original.TypeThing
)

type Action = original.Action
type Answer = original.Answer
type BaseClient = original.BaseClient
type BasicAction = original.BasicAction
type BasicAnswer = original.BasicAnswer
type BasicCreativeWork = original.BasicCreativeWork
type BasicEntitiesEntityPresentationInfo = original.BasicEntitiesEntityPresentationInfo
type BasicError = original.BasicError
type BasicGeoCoordinates = original.BasicGeoCoordinates
type BasicIdentifiable = original.BasicIdentifiable
type BasicIntangible = original.BasicIntangible
type BasicQueryContext = original.BasicQueryContext
type BasicResponse = original.BasicResponse
type BasicResponseBase = original.BasicResponseBase
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type BasicStructuredValue = original.BasicStructuredValue
type BasicThing = original.BasicThing
type CreativeWork = original.CreativeWork
type EntitiesEntityPresentationInfo = original.EntitiesEntityPresentationInfo
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type GeoCoordinates = original.GeoCoordinates
type Identifiable = original.Identifiable
type Intangible = original.Intangible
type LocalClient = original.LocalClient
type Place = original.Place
type Places = original.Places
type PostalAddress = original.PostalAddress
type QueryContext = original.QueryContext
type Response = original.Response
type ResponseBase = original.ResponseBase
type SearchAction = original.SearchAction
type SearchResponse = original.SearchResponse
type SearchResultsAnswer = original.SearchResultsAnswer
type StructuredValue = original.StructuredValue
type Thing = original.Thing

func New() BaseClient {
	return original.New()
}
func NewLocalClient() LocalClient {
	return original.NewLocalClient()
}
func NewLocalClientWithBaseURI(baseURI string) LocalClient {
	return original.NewLocalClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleEntityScenarioValues() []EntityScenario {
	return original.PossibleEntityScenarioValues()
}
func PossibleEntityTypeValues() []EntityType {
	return original.PossibleEntityTypeValues()
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return original.PossibleErrorSubCodeValues()
}
func PossibleResponseFormatValues() []ResponseFormat {
	return original.PossibleResponseFormatValues()
}
func PossibleSafeSearchValues() []SafeSearch {
	return original.PossibleSafeSearchValues()
}
func PossibleTypeBasicEntitiesEntityPresentationInfoValues() []TypeBasicEntitiesEntityPresentationInfo {
	return original.PossibleTypeBasicEntitiesEntityPresentationInfoValues()
}
func PossibleTypeBasicErrorValues() []TypeBasicError {
	return original.PossibleTypeBasicErrorValues()
}
func PossibleTypeBasicGeoCoordinatesValues() []TypeBasicGeoCoordinates {
	return original.PossibleTypeBasicGeoCoordinatesValues()
}
func PossibleTypeBasicResponseBaseValues() []TypeBasicResponseBase {
	return original.PossibleTypeBasicResponseBaseValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
