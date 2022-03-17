//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package newssearch

import original "github.com/Azure/temp/github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/newssearch"

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

type Freshness = original.Freshness

const (
	Day   Freshness = original.Day
	Month Freshness = original.Month
	Week  Freshness = original.Week
)

type SafeSearch = original.SafeSearch

const (
	Moderate SafeSearch = original.Moderate
	Off      SafeSearch = original.Off
	Strict   SafeSearch = original.Strict
)

type TextFormat = original.TextFormat

const (
	HTML TextFormat = original.HTML
	Raw  TextFormat = original.Raw
)

type Type = original.Type

const (
	TypeAnswer              Type = original.TypeAnswer
	TypeArticle             Type = original.TypeArticle
	TypeCreativeWork        Type = original.TypeCreativeWork
	TypeErrorResponse       Type = original.TypeErrorResponse
	TypeIdentifiable        Type = original.TypeIdentifiable
	TypeImageObject         Type = original.TypeImageObject
	TypeMediaObject         Type = original.TypeMediaObject
	TypeNews                Type = original.TypeNews
	TypeNewsArticle         Type = original.TypeNewsArticle
	TypeNewsTopic           Type = original.TypeNewsTopic
	TypeOrganization        Type = original.TypeOrganization
	TypeResponse            Type = original.TypeResponse
	TypeResponseBase        Type = original.TypeResponseBase
	TypeSearchResultsAnswer Type = original.TypeSearchResultsAnswer
	TypeThing               Type = original.TypeThing
	TypeTrendingTopics      Type = original.TypeTrendingTopics
	TypeVideoObject         Type = original.TypeVideoObject
)

type Answer = original.Answer
type Article = original.Article
type BaseClient = original.BaseClient
type BasicAnswer = original.BasicAnswer
type BasicArticle = original.BasicArticle
type BasicCreativeWork = original.BasicCreativeWork
type BasicIdentifiable = original.BasicIdentifiable
type BasicMediaObject = original.BasicMediaObject
type BasicResponse = original.BasicResponse
type BasicResponseBase = original.BasicResponseBase
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type BasicThing = original.BasicThing
type CreativeWork = original.CreativeWork
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type Identifiable = original.Identifiable
type ImageObject = original.ImageObject
type MediaObject = original.MediaObject
type News = original.News
type NewsArticle = original.NewsArticle
type NewsClient = original.NewsClient
type NewsTopic = original.NewsTopic
type Organization = original.Organization
type Query = original.Query
type Response = original.Response
type ResponseBase = original.ResponseBase
type SearchResultsAnswer = original.SearchResultsAnswer
type Thing = original.Thing
type TrendingTopics = original.TrendingTopics
type VideoObject = original.VideoObject

func New() BaseClient {
	return original.New()
}
func NewNewsClient() NewsClient {
	return original.NewNewsClient()
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return original.PossibleErrorSubCodeValues()
}
func PossibleFreshnessValues() []Freshness {
	return original.PossibleFreshnessValues()
}
func PossibleSafeSearchValues() []SafeSearch {
	return original.PossibleSafeSearchValues()
}
func PossibleTextFormatValues() []TextFormat {
	return original.PossibleTextFormatValues()
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
