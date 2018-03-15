// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package programmatic

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/luis/programmatic"

type AppsClient = original.AppsClient

func NewAppsClient(azureRegion AzureRegions) AppsClient {
	return original.NewAppsClient(azureRegion)
}

type BaseClient = original.BaseClient

func New(azureRegion AzureRegions) BaseClient {
	return original.New(azureRegion)
}
func NewWithoutDefaults(azureRegion AzureRegions) BaseClient {
	return original.NewWithoutDefaults(azureRegion)
}

type ExamplesClient = original.ExamplesClient

func NewExamplesClient(azureRegion AzureRegions) ExamplesClient {
	return original.NewExamplesClient(azureRegion)
}

type FeaturesClient = original.FeaturesClient

func NewFeaturesClient(azureRegion AzureRegions) FeaturesClient {
	return original.NewFeaturesClient(azureRegion)
}

type ModelClient = original.ModelClient

func NewModelClient(azureRegion AzureRegions) ModelClient {
	return original.NewModelClient(azureRegion)
}

type AzureRegions = original.AzureRegions

const (
	Australiaeast  AzureRegions = original.Australiaeast
	Brazilsouth    AzureRegions = original.Brazilsouth
	Eastasia       AzureRegions = original.Eastasia
	Eastus         AzureRegions = original.Eastus
	Eastus2        AzureRegions = original.Eastus2
	Northeurope    AzureRegions = original.Northeurope
	Southcentralus AzureRegions = original.Southcentralus
	Southeastasia  AzureRegions = original.Southeastasia
	Westcentralus  AzureRegions = original.Westcentralus
	Westeurope     AzureRegions = original.Westeurope
	Westus         AzureRegions = original.Westus
	Westus2        AzureRegions = original.Westus2
)

func PossibleAzureRegionsValues() [12]AzureRegions {
	return original.PossibleAzureRegionsValues()
}

type OperationStatusType = original.OperationStatusType

const (
	Failed  OperationStatusType = original.Failed
	FAILED  OperationStatusType = original.FAILED
	Success OperationStatusType = original.Success
)

func PossibleOperationStatusTypeValues() [3]OperationStatusType {
	return original.PossibleOperationStatusTypeValues()
}

type ReadableType = original.ReadableType

const (
	ReadableTypeClosedListEntityExtractor        ReadableType = original.ReadableTypeClosedListEntityExtractor
	ReadableTypeCompositeEntityExtractor         ReadableType = original.ReadableTypeCompositeEntityExtractor
	ReadableTypeEntityExtractor                  ReadableType = original.ReadableTypeEntityExtractor
	ReadableTypeHierarchicalChildEntityExtractor ReadableType = original.ReadableTypeHierarchicalChildEntityExtractor
	ReadableTypeHierarchicalEntityExtractor      ReadableType = original.ReadableTypeHierarchicalEntityExtractor
	ReadableTypeIntentClassifier                 ReadableType = original.ReadableTypeIntentClassifier
	ReadableTypePrebuiltEntityExtractor          ReadableType = original.ReadableTypePrebuiltEntityExtractor
)

func PossibleReadableTypeValues() [7]ReadableType {
	return original.PossibleReadableTypeValues()
}

type ReadableType1 = original.ReadableType1

const (
	ReadableType1ClosedListEntityExtractor        ReadableType1 = original.ReadableType1ClosedListEntityExtractor
	ReadableType1CompositeEntityExtractor         ReadableType1 = original.ReadableType1CompositeEntityExtractor
	ReadableType1EntityExtractor                  ReadableType1 = original.ReadableType1EntityExtractor
	ReadableType1HierarchicalChildEntityExtractor ReadableType1 = original.ReadableType1HierarchicalChildEntityExtractor
	ReadableType1HierarchicalEntityExtractor      ReadableType1 = original.ReadableType1HierarchicalEntityExtractor
	ReadableType1IntentClassifier                 ReadableType1 = original.ReadableType1IntentClassifier
	ReadableType1PrebuiltEntityExtractor          ReadableType1 = original.ReadableType1PrebuiltEntityExtractor
)

func PossibleReadableType1Values() [7]ReadableType1 {
	return original.PossibleReadableType1Values()
}

type ReadableType2 = original.ReadableType2

const (
	ReadableType2ClosedListEntityExtractor        ReadableType2 = original.ReadableType2ClosedListEntityExtractor
	ReadableType2CompositeEntityExtractor         ReadableType2 = original.ReadableType2CompositeEntityExtractor
	ReadableType2EntityExtractor                  ReadableType2 = original.ReadableType2EntityExtractor
	ReadableType2HierarchicalChildEntityExtractor ReadableType2 = original.ReadableType2HierarchicalChildEntityExtractor
	ReadableType2HierarchicalEntityExtractor      ReadableType2 = original.ReadableType2HierarchicalEntityExtractor
	ReadableType2IntentClassifier                 ReadableType2 = original.ReadableType2IntentClassifier
	ReadableType2PrebuiltEntityExtractor          ReadableType2 = original.ReadableType2PrebuiltEntityExtractor
)

func PossibleReadableType2Values() [7]ReadableType2 {
	return original.PossibleReadableType2Values()
}

type ReadableType3 = original.ReadableType3

const (
	ReadableType3ClosedListEntityExtractor        ReadableType3 = original.ReadableType3ClosedListEntityExtractor
	ReadableType3CompositeEntityExtractor         ReadableType3 = original.ReadableType3CompositeEntityExtractor
	ReadableType3EntityExtractor                  ReadableType3 = original.ReadableType3EntityExtractor
	ReadableType3HierarchicalChildEntityExtractor ReadableType3 = original.ReadableType3HierarchicalChildEntityExtractor
	ReadableType3HierarchicalEntityExtractor      ReadableType3 = original.ReadableType3HierarchicalEntityExtractor
	ReadableType3IntentClassifier                 ReadableType3 = original.ReadableType3IntentClassifier
	ReadableType3PrebuiltEntityExtractor          ReadableType3 = original.ReadableType3PrebuiltEntityExtractor
)

func PossibleReadableType3Values() [7]ReadableType3 {
	return original.PossibleReadableType3Values()
}

type Status = original.Status

const (
	StatusFail       Status = original.StatusFail
	StatusInProgress Status = original.StatusInProgress
	StatusQueued     Status = original.StatusQueued
	StatusSuccess    Status = original.StatusSuccess
	StatusUpToDate   Status = original.StatusUpToDate
)

func PossibleStatusValues() [5]Status {
	return original.PossibleStatusValues()
}

type Status1 = original.Status1

const (
	Status1Fail       Status1 = original.Status1Fail
	Status1InProgress Status1 = original.Status1InProgress
	Status1Queued     Status1 = original.Status1Queued
	Status1Success    Status1 = original.Status1Success
	Status1UpToDate   Status1 = original.Status1UpToDate
)

func PossibleStatus1Values() [5]Status1 {
	return original.PossibleStatus1Values()
}

type TrainingStatus = original.TrainingStatus

const (
	InProgress    TrainingStatus = original.InProgress
	NeedsTraining TrainingStatus = original.NeedsTraining
	Trained       TrainingStatus = original.Trained
)

func PossibleTrainingStatusValues() [3]TrainingStatus {
	return original.PossibleTrainingStatusValues()
}

type ApplicationCreateObject = original.ApplicationCreateObject
type ApplicationInfoResponse = original.ApplicationInfoResponse
type ApplicationPublishObject = original.ApplicationPublishObject
type ApplicationSettings = original.ApplicationSettings
type ApplicationSettingUpdateObject = original.ApplicationSettingUpdateObject
type ApplicationUpdateObject = original.ApplicationUpdateObject
type AvailableCulture = original.AvailableCulture
type AvailablePrebuiltEntityModel = original.AvailablePrebuiltEntityModel
type BatchLabelExample = original.BatchLabelExample
type ChildEntity = original.ChildEntity
type ClosedList = original.ClosedList
type ClosedListEntityExtractor = original.ClosedListEntityExtractor
type ClosedListModelCreateObject = original.ClosedListModelCreateObject
type ClosedListModelPatchObject = original.ClosedListModelPatchObject
type ClosedListModelUpdateObject = original.ClosedListModelUpdateObject
type CollaboratorsArray = original.CollaboratorsArray
type CompositeChildModelCreateObject = original.CompositeChildModelCreateObject
type CompositeEntityExtractor = original.CompositeEntityExtractor
type CompositeEntityModel = original.CompositeEntityModel
type CustomPrebuiltModel = original.CustomPrebuiltModel
type EndpointInfo = original.EndpointInfo
type EnqueueTrainingResponse = original.EnqueueTrainingResponse
type EntitiesSuggestionExample = original.EntitiesSuggestionExample
type EntityExtractor = original.EntityExtractor
type EntityLabel = original.EntityLabel
type EntityLabelObject = original.EntityLabelObject
type EntityPrediction = original.EntityPrediction
type ErrorResponse = original.ErrorResponse
type ExampleLabelObject = original.ExampleLabelObject
type FeatureInfoObject = original.FeatureInfoObject
type FeaturesResponseObject = original.FeaturesResponseObject
type HierarchicalChildEntity = original.HierarchicalChildEntity
type HierarchicalChildModelCreateObject = original.HierarchicalChildModelCreateObject
type HierarchicalChildModelUpdateObject = original.HierarchicalChildModelUpdateObject
type HierarchicalEntityExtractor = original.HierarchicalEntityExtractor
type HierarchicalEntityModel = original.HierarchicalEntityModel
type HierarchicalModel = original.HierarchicalModel
type Int32 = original.Int32
type IntentClassifier = original.IntentClassifier
type IntentPrediction = original.IntentPrediction
type IntentsSuggestionExample = original.IntentsSuggestionExample
type JSONEntity = original.JSONEntity
type JSONModelFeature = original.JSONModelFeature
type JSONRegexFeature = original.JSONRegexFeature
type JSONUtterance = original.JSONUtterance
type LabeledUtterance = original.LabeledUtterance
type LabelExampleResponse = original.LabelExampleResponse
type ListApplicationInfoResponse = original.ListApplicationInfoResponse
type ListAvailableCulture = original.ListAvailableCulture
type ListAvailablePrebuiltEntityModel = original.ListAvailablePrebuiltEntityModel
type ListBatchLabelExample = original.ListBatchLabelExample
type ListClosedListEntityExtractor = original.ListClosedListEntityExtractor
type ListCompositeEntityExtractor = original.ListCompositeEntityExtractor
type ListCustomPrebuiltModel = original.ListCustomPrebuiltModel
type ListEntitiesSuggestionExample = original.ListEntitiesSuggestionExample
type ListEntityExtractor = original.ListEntityExtractor
type ListHierarchicalEntityExtractor = original.ListHierarchicalEntityExtractor
type ListIntentClassifier = original.ListIntentClassifier
type ListIntentsSuggestionExample = original.ListIntentsSuggestionExample
type ListLabeledUtterance = original.ListLabeledUtterance
type ListModelInfoResponse = original.ListModelInfoResponse
type ListModelTrainingInfo = original.ListModelTrainingInfo
type ListPhraseListFeatureInfo = original.ListPhraseListFeatureInfo
type ListPrebuiltDomain = original.ListPrebuiltDomain
type ListPrebuiltEntityExtractor = original.ListPrebuiltEntityExtractor
type ListString = original.ListString
type ListUUID = original.ListUUID
type ListVersionInfo = original.ListVersionInfo
type LuisApp = original.LuisApp
type ModelCreateObject = original.ModelCreateObject
type ModelInfo = original.ModelInfo
type ModelInfoResponse = original.ModelInfoResponse
type ModelTrainingDetails = original.ModelTrainingDetails
type ModelTrainingInfo = original.ModelTrainingInfo
type ModelUpdateObject = original.ModelUpdateObject
type OperationError = original.OperationError
type OperationStatus = original.OperationStatus
type PatternCreateObject = original.PatternCreateObject
type PatternFeatureInfo = original.PatternFeatureInfo
type PatternUpdateObject = original.PatternUpdateObject
type PersonalAssistantsResponse = original.PersonalAssistantsResponse
type PhraselistCreateObject = original.PhraselistCreateObject
type PhraseListFeatureInfo = original.PhraseListFeatureInfo
type PhraselistUpdateObject = original.PhraselistUpdateObject
type PrebuiltDomain = original.PrebuiltDomain
type PrebuiltDomainCreateBaseObject = original.PrebuiltDomainCreateBaseObject
type PrebuiltDomainCreateObject = original.PrebuiltDomainCreateObject
type PrebuiltDomainItem = original.PrebuiltDomainItem
type PrebuiltDomainModelCreateObject = original.PrebuiltDomainModelCreateObject
type PrebuiltEntityExtractor = original.PrebuiltEntityExtractor
type ProductionOrStagingEndpointInfo = original.ProductionOrStagingEndpointInfo
type ReadCloser = original.ReadCloser
type SetString = original.SetString
type String = original.String
type SubClosedList = original.SubClosedList
type SubClosedListResponse = original.SubClosedListResponse
type TaskUpdateObject = original.TaskUpdateObject
type UserAccessList = original.UserAccessList
type UserCollaborator = original.UserCollaborator
type UUID = original.UUID
type VersionInfo = original.VersionInfo
type WordListBaseUpdateObject = original.WordListBaseUpdateObject
type WordListObject = original.WordListObject
type PermissionsClient = original.PermissionsClient

func NewPermissionsClient(azureRegion AzureRegions) PermissionsClient {
	return original.NewPermissionsClient(azureRegion)
}

type TrainClient = original.TrainClient

func NewTrainClient(azureRegion AzureRegions) TrainClient {
	return original.NewTrainClient(azureRegion)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

type VersionsClient = original.VersionsClient

func NewVersionsClient(azureRegion AzureRegions) VersionsClient {
	return original.NewVersionsClient(azureRegion)
}
