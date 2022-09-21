//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package entitysearch

import original "github.com/Azure/dev/azure-sdk-for-go/services/cognitiveservices/v1.0/entitysearch"

const (
	DefaultEndpoint = original.DefaultEndpoint
)

type AnswerType = original.AnswerType

const (
	AnswerTypeEntities AnswerType = original.AnswerTypeEntities
	AnswerTypePlaces   AnswerType = original.AnswerTypePlaces
)

type EntityQueryScenario = original.EntityQueryScenario

const (
	Disambiguation                   EntityQueryScenario = original.Disambiguation
	DominantEntity                   EntityQueryScenario = original.DominantEntity
	DominantEntityWithDisambiguation EntityQueryScenario = original.DominantEntityWithDisambiguation
	List                             EntityQueryScenario = original.List
	ListWithPivot                    EntityQueryScenario = original.ListWithPivot
)

type EntityScenario = original.EntityScenario

const (
	EntityScenarioDisambiguationItem EntityScenario = original.EntityScenarioDisambiguationItem
	EntityScenarioDominantEntity     EntityScenario = original.EntityScenarioDominantEntity
	EntityScenarioListItem           EntityScenario = original.EntityScenarioListItem
)

type EntityType = original.EntityType

const (
	EntityTypeActor               EntityType = original.EntityTypeActor
	EntityTypeAnimal              EntityType = original.EntityTypeAnimal
	EntityTypeArtist              EntityType = original.EntityTypeArtist
	EntityTypeAttorney            EntityType = original.EntityTypeAttorney
	EntityTypeAttraction          EntityType = original.EntityTypeAttraction
	EntityTypeBook                EntityType = original.EntityTypeBook
	EntityTypeCar                 EntityType = original.EntityTypeCar
	EntityTypeCity                EntityType = original.EntityTypeCity
	EntityTypeCollegeOrUniversity EntityType = original.EntityTypeCollegeOrUniversity
	EntityTypeComposition         EntityType = original.EntityTypeComposition
	EntityTypeContinent           EntityType = original.EntityTypeContinent
	EntityTypeCountry             EntityType = original.EntityTypeCountry
	EntityTypeDrug                EntityType = original.EntityTypeDrug
	EntityTypeEvent               EntityType = original.EntityTypeEvent
	EntityTypeFood                EntityType = original.EntityTypeFood
	EntityTypeGeneric             EntityType = original.EntityTypeGeneric
	EntityTypeHotel               EntityType = original.EntityTypeHotel
	EntityTypeHouse               EntityType = original.EntityTypeHouse
	EntityTypeLocalBusiness       EntityType = original.EntityTypeLocalBusiness
	EntityTypeLocality            EntityType = original.EntityTypeLocality
	EntityTypeMedia               EntityType = original.EntityTypeMedia
	EntityTypeMinorRegion         EntityType = original.EntityTypeMinorRegion
	EntityTypeMovie               EntityType = original.EntityTypeMovie
	EntityTypeMusicAlbum          EntityType = original.EntityTypeMusicAlbum
	EntityTypeMusicGroup          EntityType = original.EntityTypeMusicGroup
	EntityTypeMusicRecording      EntityType = original.EntityTypeMusicRecording
	EntityTypeNeighborhood        EntityType = original.EntityTypeNeighborhood
	EntityTypeOrganization        EntityType = original.EntityTypeOrganization
	EntityTypeOther               EntityType = original.EntityTypeOther
	EntityTypePerson              EntityType = original.EntityTypePerson
	EntityTypePlace               EntityType = original.EntityTypePlace
	EntityTypePointOfInterest     EntityType = original.EntityTypePointOfInterest
	EntityTypePostalCode          EntityType = original.EntityTypePostalCode
	EntityTypeProduct             EntityType = original.EntityTypeProduct
	EntityTypeRadioStation        EntityType = original.EntityTypeRadioStation
	EntityTypeRegion              EntityType = original.EntityTypeRegion
	EntityTypeRestaurant          EntityType = original.EntityTypeRestaurant
	EntityTypeSchool              EntityType = original.EntityTypeSchool
	EntityTypeSpeciality          EntityType = original.EntityTypeSpeciality
	EntityTypeSportsTeam          EntityType = original.EntityTypeSportsTeam
	EntityTypeState               EntityType = original.EntityTypeState
	EntityTypeStreetAddress       EntityType = original.EntityTypeStreetAddress
	EntityTypeSubRegion           EntityType = original.EntityTypeSubRegion
	EntityTypeTelevisionSeason    EntityType = original.EntityTypeTelevisionSeason
	EntityTypeTelevisionShow      EntityType = original.EntityTypeTelevisionShow
	EntityTypeTheaterPlay         EntityType = original.EntityTypeTheaterPlay
	EntityTypeTouristAttraction   EntityType = original.EntityTypeTouristAttraction
	EntityTypeTravel              EntityType = original.EntityTypeTravel
	EntityTypeVideoGame           EntityType = original.EntityTypeVideoGame
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
	TypeContractualRulesAttribution        Type = original.TypeContractualRulesAttribution
	TypeContractualRulesContractualRule    Type = original.TypeContractualRulesContractualRule
	TypeContractualRulesLicenseAttribution Type = original.TypeContractualRulesLicenseAttribution
	TypeContractualRulesLinkAttribution    Type = original.TypeContractualRulesLinkAttribution
	TypeContractualRulesMediaAttribution   Type = original.TypeContractualRulesMediaAttribution
	TypeContractualRulesTextAttribution    Type = original.TypeContractualRulesTextAttribution
)

type TypeBasicResponseBase = original.TypeBasicResponseBase

const (
	TypeAirport               TypeBasicResponseBase = original.TypeAirport
	TypeAnswer                TypeBasicResponseBase = original.TypeAnswer
	TypeCivicStructure        TypeBasicResponseBase = original.TypeCivicStructure
	TypeCreativeWork          TypeBasicResponseBase = original.TypeCreativeWork
	TypeEntertainmentBusiness TypeBasicResponseBase = original.TypeEntertainmentBusiness
	TypeEntities              TypeBasicResponseBase = original.TypeEntities
	TypeErrorResponse         TypeBasicResponseBase = original.TypeErrorResponse
	TypeFoodEstablishment     TypeBasicResponseBase = original.TypeFoodEstablishment
	TypeHotel                 TypeBasicResponseBase = original.TypeHotel
	TypeIdentifiable          TypeBasicResponseBase = original.TypeIdentifiable
	TypeImageObject           TypeBasicResponseBase = original.TypeImageObject
	TypeIntangible            TypeBasicResponseBase = original.TypeIntangible
	TypeLicense               TypeBasicResponseBase = original.TypeLicense
	TypeLocalBusiness         TypeBasicResponseBase = original.TypeLocalBusiness
	TypeLodgingBusiness       TypeBasicResponseBase = original.TypeLodgingBusiness
	TypeMediaObject           TypeBasicResponseBase = original.TypeMediaObject
	TypeMovieTheater          TypeBasicResponseBase = original.TypeMovieTheater
	TypeOrganization          TypeBasicResponseBase = original.TypeOrganization
	TypePlace                 TypeBasicResponseBase = original.TypePlace
	TypePlaces                TypeBasicResponseBase = original.TypePlaces
	TypePostalAddress         TypeBasicResponseBase = original.TypePostalAddress
	TypeResponse              TypeBasicResponseBase = original.TypeResponse
	TypeResponseBase          TypeBasicResponseBase = original.TypeResponseBase
	TypeRestaurant            TypeBasicResponseBase = original.TypeRestaurant
	TypeSearchResponse        TypeBasicResponseBase = original.TypeSearchResponse
	TypeSearchResultsAnswer   TypeBasicResponseBase = original.TypeSearchResultsAnswer
	TypeStructuredValue       TypeBasicResponseBase = original.TypeStructuredValue
	TypeThing                 TypeBasicResponseBase = original.TypeThing
	TypeTouristAttraction     TypeBasicResponseBase = original.TypeTouristAttraction
)

type Airport = original.Airport
type Answer = original.Answer
type BaseClient = original.BaseClient
type BasicAnswer = original.BasicAnswer
type BasicCivicStructure = original.BasicCivicStructure
type BasicContractualRulesAttribution = original.BasicContractualRulesAttribution
type BasicContractualRulesContractualRule = original.BasicContractualRulesContractualRule
type BasicCreativeWork = original.BasicCreativeWork
type BasicEntertainmentBusiness = original.BasicEntertainmentBusiness
type BasicFoodEstablishment = original.BasicFoodEstablishment
type BasicIdentifiable = original.BasicIdentifiable
type BasicIntangible = original.BasicIntangible
type BasicLocalBusiness = original.BasicLocalBusiness
type BasicLodgingBusiness = original.BasicLodgingBusiness
type BasicMediaObject = original.BasicMediaObject
type BasicPlace = original.BasicPlace
type BasicResponse = original.BasicResponse
type BasicResponseBase = original.BasicResponseBase
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type BasicStructuredValue = original.BasicStructuredValue
type BasicThing = original.BasicThing
type CivicStructure = original.CivicStructure
type ContractualRulesAttribution = original.ContractualRulesAttribution
type ContractualRulesContractualRule = original.ContractualRulesContractualRule
type ContractualRulesLicenseAttribution = original.ContractualRulesLicenseAttribution
type ContractualRulesLinkAttribution = original.ContractualRulesLinkAttribution
type ContractualRulesMediaAttribution = original.ContractualRulesMediaAttribution
type ContractualRulesTextAttribution = original.ContractualRulesTextAttribution
type CreativeWork = original.CreativeWork
type EntertainmentBusiness = original.EntertainmentBusiness
type Entities = original.Entities
type EntitiesClient = original.EntitiesClient
type EntitiesEntityPresentationInfo = original.EntitiesEntityPresentationInfo
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type FoodEstablishment = original.FoodEstablishment
type Hotel = original.Hotel
type Identifiable = original.Identifiable
type ImageObject = original.ImageObject
type Intangible = original.Intangible
type License = original.License
type LocalBusiness = original.LocalBusiness
type LodgingBusiness = original.LodgingBusiness
type MediaObject = original.MediaObject
type MovieTheater = original.MovieTheater
type Organization = original.Organization
type Place = original.Place
type Places = original.Places
type PostalAddress = original.PostalAddress
type QueryContext = original.QueryContext
type Response = original.Response
type ResponseBase = original.ResponseBase
type Restaurant = original.Restaurant
type SearchResponse = original.SearchResponse
type SearchResultsAnswer = original.SearchResultsAnswer
type StructuredValue = original.StructuredValue
type Thing = original.Thing
type TouristAttraction = original.TouristAttraction

func New() BaseClient {
	return original.New()
}
func NewEntitiesClient() EntitiesClient {
	return original.NewEntitiesClient()
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleAnswerTypeValues() []AnswerType {
	return original.PossibleAnswerTypeValues()
}
func PossibleEntityQueryScenarioValues() []EntityQueryScenario {
	return original.PossibleEntityQueryScenarioValues()
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
