// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package speechservices

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/SpeechServices"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DataImportKind = original.DataImportKind

const (
	Acoustic           DataImportKind = original.Acoustic
	CustomVoice        DataImportKind = original.CustomVoice
	Language           DataImportKind = original.Language
	LanguageGeneration DataImportKind = original.LanguageGeneration
	None               DataImportKind = original.None
	Pronunciation      DataImportKind = original.Pronunciation
)

type DataImportKind1 = original.DataImportKind1

const (
	DataImportKind1Acoustic           DataImportKind1 = original.DataImportKind1Acoustic
	DataImportKind1CustomVoice        DataImportKind1 = original.DataImportKind1CustomVoice
	DataImportKind1Language           DataImportKind1 = original.DataImportKind1Language
	DataImportKind1LanguageGeneration DataImportKind1 = original.DataImportKind1LanguageGeneration
	DataImportKind1None               DataImportKind1 = original.DataImportKind1None
	DataImportKind1Pronunciation      DataImportKind1 = original.DataImportKind1Pronunciation
)

type EndpointKind = original.EndpointKind

const (
	EndpointKindCustomVoice            EndpointKind = original.EndpointKindCustomVoice
	EndpointKindLanguageGeneration     EndpointKind = original.EndpointKindLanguageGeneration
	EndpointKindLanguageIdentification EndpointKind = original.EndpointKindLanguageIdentification
	EndpointKindNone                   EndpointKind = original.EndpointKindNone
	EndpointKindSpeechRecognition      EndpointKind = original.EndpointKindSpeechRecognition
)

type ModelKind = original.ModelKind

const (
	ModelKindAcoustic               ModelKind = original.ModelKindAcoustic
	ModelKindAcousticAndLanguage    ModelKind = original.ModelKindAcousticAndLanguage
	ModelKindCustomVoice            ModelKind = original.ModelKindCustomVoice
	ModelKindLanguage               ModelKind = original.ModelKindLanguage
	ModelKindLanguageGeneration     ModelKind = original.ModelKindLanguageGeneration
	ModelKindLanguageIdentification ModelKind = original.ModelKindLanguageIdentification
	ModelKindNone                   ModelKind = original.ModelKindNone
	ModelKindSentiment              ModelKind = original.ModelKindSentiment
)

type ModelKind1 = original.ModelKind1

const (
	ModelKind1Acoustic               ModelKind1 = original.ModelKind1Acoustic
	ModelKind1AcousticAndLanguage    ModelKind1 = original.ModelKind1AcousticAndLanguage
	ModelKind1CustomVoice            ModelKind1 = original.ModelKind1CustomVoice
	ModelKind1Language               ModelKind1 = original.ModelKind1Language
	ModelKind1LanguageGeneration     ModelKind1 = original.ModelKind1LanguageGeneration
	ModelKind1LanguageIdentification ModelKind1 = original.ModelKind1LanguageIdentification
	ModelKind1None                   ModelKind1 = original.ModelKind1None
	ModelKind1Sentiment              ModelKind1 = original.ModelKind1Sentiment
)

type ModelKind2 = original.ModelKind2

const (
	ModelKind2Acoustic               ModelKind2 = original.ModelKind2Acoustic
	ModelKind2AcousticAndLanguage    ModelKind2 = original.ModelKind2AcousticAndLanguage
	ModelKind2CustomVoice            ModelKind2 = original.ModelKind2CustomVoice
	ModelKind2Language               ModelKind2 = original.ModelKind2Language
	ModelKind2LanguageGeneration     ModelKind2 = original.ModelKind2LanguageGeneration
	ModelKind2LanguageIdentification ModelKind2 = original.ModelKind2LanguageIdentification
	ModelKind2None                   ModelKind2 = original.ModelKind2None
	ModelKind2Sentiment              ModelKind2 = original.ModelKind2Sentiment
)

type Status = original.Status

const (
	Failed     Status = original.Failed
	NotStarted Status = original.NotStarted
	Running    Status = original.Running
	Succeeded  Status = original.Succeeded
)

type Status1 = original.Status1

const (
	Status1Failed     Status1 = original.Status1Failed
	Status1NotStarted Status1 = original.Status1NotStarted
	Status1Running    Status1 = original.Status1Running
	Status1Succeeded  Status1 = original.Status1Succeeded
)

type Status2 = original.Status2

const (
	Status2Failed     Status2 = original.Status2Failed
	Status2NotStarted Status2 = original.Status2NotStarted
	Status2Running    Status2 = original.Status2Running
	Status2Succeeded  Status2 = original.Status2Succeeded
)

type Status3 = original.Status3

const (
	Status3Failed     Status3 = original.Status3Failed
	Status3NotStarted Status3 = original.Status3NotStarted
	Status3Running    Status3 = original.Status3Running
	Status3Succeeded  Status3 = original.Status3Succeeded
)

type Status4 = original.Status4

const (
	Status4Failed     Status4 = original.Status4Failed
	Status4NotStarted Status4 = original.Status4NotStarted
	Status4Running    Status4 = original.Status4Running
	Status4Succeeded  Status4 = original.Status4Succeeded
)

type Status5 = original.Status5

const (
	Status5Failed     Status5 = original.Status5Failed
	Status5NotStarted Status5 = original.Status5NotStarted
	Status5Running    Status5 = original.Status5Running
	Status5Succeeded  Status5 = original.Status5Succeeded
)

type Status6 = original.Status6

const (
	Error   Status6 = original.Error
	Healthy Status6 = original.Healthy
	Sick    Status6 = original.Sick
)

type Status7 = original.Status7

const (
	Status7Error   Status7 = original.Status7Error
	Status7Healthy Status7 = original.Status7Healthy
	Status7Sick    Status7 = original.Status7Sick
)

type Status8 = original.Status8

const (
	Status8Failed     Status8 = original.Status8Failed
	Status8NotStarted Status8 = original.Status8NotStarted
	Status8Running    Status8 = original.Status8Running
	Status8Succeeded  Status8 = original.Status8Succeeded
)

type VoiceTestKind = original.VoiceTestKind

const (
	VoiceTestKindNone VoiceTestKind = original.VoiceTestKindNone
	VoiceTestKindSSML VoiceTestKind = original.VoiceTestKindSSML
	VoiceTestKindText VoiceTestKind = original.VoiceTestKindText
)

type VoiceTestKind1 = original.VoiceTestKind1

const (
	VoiceTestKind1None VoiceTestKind1 = original.VoiceTestKind1None
	VoiceTestKind1SSML VoiceTestKind1 = original.VoiceTestKind1SSML
	VoiceTestKind1Text VoiceTestKind1 = original.VoiceTestKind1Text
)

type AccuracyTestClient = original.AccuracyTestClient
type AccuracyTestCreateFuture = original.AccuracyTestCreateFuture
type AccuracyTestsClient = original.AccuracyTestsClient
type BaseClient = original.BaseClient
type Component = original.Component
type Dataset = original.Dataset
type DatasetClient = original.DatasetClient
type DatasetDefinition = original.DatasetDefinition
type DatasetIdentity = original.DatasetIdentity
type DatasetUpdate = original.DatasetUpdate
type DatasetUploadFuture = original.DatasetUploadFuture
type DatasetsClient = original.DatasetsClient
type DeploymentClient = original.DeploymentClient
type Endpoint = original.Endpoint
type EndpointClient = original.EndpointClient
type EndpointCreateFuture = original.EndpointCreateFuture
type EndpointData = original.EndpointData
type EndpointDataClient = original.EndpointDataClient
type EndpointDataDefinition = original.EndpointDataDefinition
type EndpointDataExportClient = original.EndpointDataExportClient
type EndpointDataExportCreateFuture = original.EndpointDataExportCreateFuture
type EndpointDataExportsClient = original.EndpointDataExportsClient
type EndpointDefinition = original.EndpointDefinition
type EndpointUpdate = original.EndpointUpdate
type EndpointsClient = original.EndpointsClient
type ErrorContent = original.ErrorContent
type ErrorDetail = original.ErrorDetail
type HealthStatusClient = original.HealthStatusClient
type HealthStatusResponse = original.HealthStatusResponse
type IReadOnlyDictionary = original.IReadOnlyDictionary
type InnerError = original.InnerError
type LanguageGenerationEndpointClient = original.LanguageGenerationEndpointClient
type LanguageGenerationEndpointCreateFuture = original.LanguageGenerationEndpointCreateFuture
type LanguageGenerationEndpointsClient = original.LanguageGenerationEndpointsClient
type LanguageGenerationModelClient = original.LanguageGenerationModelClient
type LanguageGenerationModelCreateFuture = original.LanguageGenerationModelCreateFuture
type LanguageGenerationModelsClient = original.LanguageGenerationModelsClient
type Model = original.Model
type ModelClient = original.ModelClient
type ModelCreateFuture = original.ModelCreateFuture
type ModelDefinition = original.ModelDefinition
type ModelIdentity = original.ModelIdentity
type ModelUpdate = original.ModelUpdate
type ModelsClient = original.ModelsClient
type SetObject = original.SetObject
type SpeechEndpointDefinition = original.SpeechEndpointDefinition
type SpeechModelDefinition = original.SpeechModelDefinition
type SupportedLocalesForDatasetsClient = original.SupportedLocalesForDatasetsClient
type SupportedLocalesForEndpointsClient = original.SupportedLocalesForEndpointsClient
type SupportedLocalesForLanguageGenerationEndpointsClient = original.SupportedLocalesForLanguageGenerationEndpointsClient
type SupportedLocalesForLanguageGenerationModelsClient = original.SupportedLocalesForLanguageGenerationModelsClient
type SupportedLocalesForModelsClient = original.SupportedLocalesForModelsClient
type SupportedLocalesForTranscriptionsClient = original.SupportedLocalesForTranscriptionsClient
type SupportedLocalesForVoiceDatasetsClient = original.SupportedLocalesForVoiceDatasetsClient
type SupportedLocalesForVoiceEndpointsClient = original.SupportedLocalesForVoiceEndpointsClient
type SupportedLocalesForVoiceModelsClient = original.SupportedLocalesForVoiceModelsClient
type Test = original.Test
type TestDefinition = original.TestDefinition
type TestUpdate = original.TestUpdate
type Transcription = original.Transcription
type TranscriptionClient = original.TranscriptionClient
type TranscriptionDefinition = original.TranscriptionDefinition
type TranscriptionUpdate = original.TranscriptionUpdate
type TranscriptionsClient = original.TranscriptionsClient
type VoiceDatasetClient = original.VoiceDatasetClient
type VoiceDatasetUploadFuture = original.VoiceDatasetUploadFuture
type VoiceDatasetsClient = original.VoiceDatasetsClient
type VoiceDeploymentClient = original.VoiceDeploymentClient
type VoiceDeploymentCreateFuture = original.VoiceDeploymentCreateFuture
type VoiceDeploymentsClient = original.VoiceDeploymentsClient
type VoiceEndpointClient = original.VoiceEndpointClient
type VoiceModelClient = original.VoiceModelClient
type VoiceModelCreateFuture = original.VoiceModelCreateFuture
type VoiceModelsClient = original.VoiceModelsClient
type VoiceTest = original.VoiceTest
type VoiceTestClient = original.VoiceTestClient
type VoiceTestCreateFuture = original.VoiceTestCreateFuture
type VoiceTestDefinition = original.VoiceTestDefinition
type WebHookConfigurationSecret = original.WebHookConfigurationSecret
type WebHookUpdate = original.WebHookUpdate

func New() BaseClient {
	return original.New()
}
func NewAccuracyTestClient() AccuracyTestClient {
	return original.NewAccuracyTestClient()
}
func NewAccuracyTestClientWithBaseURI(baseURI string) AccuracyTestClient {
	return original.NewAccuracyTestClientWithBaseURI(baseURI)
}
func NewAccuracyTestsClient() AccuracyTestsClient {
	return original.NewAccuracyTestsClient()
}
func NewAccuracyTestsClientWithBaseURI(baseURI string) AccuracyTestsClient {
	return original.NewAccuracyTestsClientWithBaseURI(baseURI)
}
func NewDatasetClient() DatasetClient {
	return original.NewDatasetClient()
}
func NewDatasetClientWithBaseURI(baseURI string) DatasetClient {
	return original.NewDatasetClientWithBaseURI(baseURI)
}
func NewDatasetsClient() DatasetsClient {
	return original.NewDatasetsClient()
}
func NewDatasetsClientWithBaseURI(baseURI string) DatasetsClient {
	return original.NewDatasetsClientWithBaseURI(baseURI)
}
func NewDeploymentClient() DeploymentClient {
	return original.NewDeploymentClient()
}
func NewDeploymentClientWithBaseURI(baseURI string) DeploymentClient {
	return original.NewDeploymentClientWithBaseURI(baseURI)
}
func NewEndpointClient() EndpointClient {
	return original.NewEndpointClient()
}
func NewEndpointClientWithBaseURI(baseURI string) EndpointClient {
	return original.NewEndpointClientWithBaseURI(baseURI)
}
func NewEndpointDataClient() EndpointDataClient {
	return original.NewEndpointDataClient()
}
func NewEndpointDataClientWithBaseURI(baseURI string) EndpointDataClient {
	return original.NewEndpointDataClientWithBaseURI(baseURI)
}
func NewEndpointDataExportClient() EndpointDataExportClient {
	return original.NewEndpointDataExportClient()
}
func NewEndpointDataExportClientWithBaseURI(baseURI string) EndpointDataExportClient {
	return original.NewEndpointDataExportClientWithBaseURI(baseURI)
}
func NewEndpointDataExportsClient() EndpointDataExportsClient {
	return original.NewEndpointDataExportsClient()
}
func NewEndpointDataExportsClientWithBaseURI(baseURI string) EndpointDataExportsClient {
	return original.NewEndpointDataExportsClientWithBaseURI(baseURI)
}
func NewEndpointsClient() EndpointsClient {
	return original.NewEndpointsClient()
}
func NewEndpointsClientWithBaseURI(baseURI string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI)
}
func NewHealthStatusClient() HealthStatusClient {
	return original.NewHealthStatusClient()
}
func NewHealthStatusClientWithBaseURI(baseURI string) HealthStatusClient {
	return original.NewHealthStatusClientWithBaseURI(baseURI)
}
func NewLanguageGenerationEndpointClient() LanguageGenerationEndpointClient {
	return original.NewLanguageGenerationEndpointClient()
}
func NewLanguageGenerationEndpointClientWithBaseURI(baseURI string) LanguageGenerationEndpointClient {
	return original.NewLanguageGenerationEndpointClientWithBaseURI(baseURI)
}
func NewLanguageGenerationEndpointsClient() LanguageGenerationEndpointsClient {
	return original.NewLanguageGenerationEndpointsClient()
}
func NewLanguageGenerationEndpointsClientWithBaseURI(baseURI string) LanguageGenerationEndpointsClient {
	return original.NewLanguageGenerationEndpointsClientWithBaseURI(baseURI)
}
func NewLanguageGenerationModelClient() LanguageGenerationModelClient {
	return original.NewLanguageGenerationModelClient()
}
func NewLanguageGenerationModelClientWithBaseURI(baseURI string) LanguageGenerationModelClient {
	return original.NewLanguageGenerationModelClientWithBaseURI(baseURI)
}
func NewLanguageGenerationModelsClient() LanguageGenerationModelsClient {
	return original.NewLanguageGenerationModelsClient()
}
func NewLanguageGenerationModelsClientWithBaseURI(baseURI string) LanguageGenerationModelsClient {
	return original.NewLanguageGenerationModelsClientWithBaseURI(baseURI)
}
func NewModelClient() ModelClient {
	return original.NewModelClient()
}
func NewModelClientWithBaseURI(baseURI string) ModelClient {
	return original.NewModelClientWithBaseURI(baseURI)
}
func NewModelsClient() ModelsClient {
	return original.NewModelsClient()
}
func NewModelsClientWithBaseURI(baseURI string) ModelsClient {
	return original.NewModelsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForDatasetsClient() SupportedLocalesForDatasetsClient {
	return original.NewSupportedLocalesForDatasetsClient()
}
func NewSupportedLocalesForDatasetsClientWithBaseURI(baseURI string) SupportedLocalesForDatasetsClient {
	return original.NewSupportedLocalesForDatasetsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForEndpointsClient() SupportedLocalesForEndpointsClient {
	return original.NewSupportedLocalesForEndpointsClient()
}
func NewSupportedLocalesForEndpointsClientWithBaseURI(baseURI string) SupportedLocalesForEndpointsClient {
	return original.NewSupportedLocalesForEndpointsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForLanguageGenerationEndpointsClient() SupportedLocalesForLanguageGenerationEndpointsClient {
	return original.NewSupportedLocalesForLanguageGenerationEndpointsClient()
}
func NewSupportedLocalesForLanguageGenerationEndpointsClientWithBaseURI(baseURI string) SupportedLocalesForLanguageGenerationEndpointsClient {
	return original.NewSupportedLocalesForLanguageGenerationEndpointsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForLanguageGenerationModelsClient() SupportedLocalesForLanguageGenerationModelsClient {
	return original.NewSupportedLocalesForLanguageGenerationModelsClient()
}
func NewSupportedLocalesForLanguageGenerationModelsClientWithBaseURI(baseURI string) SupportedLocalesForLanguageGenerationModelsClient {
	return original.NewSupportedLocalesForLanguageGenerationModelsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForModelsClient() SupportedLocalesForModelsClient {
	return original.NewSupportedLocalesForModelsClient()
}
func NewSupportedLocalesForModelsClientWithBaseURI(baseURI string) SupportedLocalesForModelsClient {
	return original.NewSupportedLocalesForModelsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForTranscriptionsClient() SupportedLocalesForTranscriptionsClient {
	return original.NewSupportedLocalesForTranscriptionsClient()
}
func NewSupportedLocalesForTranscriptionsClientWithBaseURI(baseURI string) SupportedLocalesForTranscriptionsClient {
	return original.NewSupportedLocalesForTranscriptionsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForVoiceDatasetsClient() SupportedLocalesForVoiceDatasetsClient {
	return original.NewSupportedLocalesForVoiceDatasetsClient()
}
func NewSupportedLocalesForVoiceDatasetsClientWithBaseURI(baseURI string) SupportedLocalesForVoiceDatasetsClient {
	return original.NewSupportedLocalesForVoiceDatasetsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForVoiceEndpointsClient() SupportedLocalesForVoiceEndpointsClient {
	return original.NewSupportedLocalesForVoiceEndpointsClient()
}
func NewSupportedLocalesForVoiceEndpointsClientWithBaseURI(baseURI string) SupportedLocalesForVoiceEndpointsClient {
	return original.NewSupportedLocalesForVoiceEndpointsClientWithBaseURI(baseURI)
}
func NewSupportedLocalesForVoiceModelsClient() SupportedLocalesForVoiceModelsClient {
	return original.NewSupportedLocalesForVoiceModelsClient()
}
func NewSupportedLocalesForVoiceModelsClientWithBaseURI(baseURI string) SupportedLocalesForVoiceModelsClient {
	return original.NewSupportedLocalesForVoiceModelsClientWithBaseURI(baseURI)
}
func NewTranscriptionClient() TranscriptionClient {
	return original.NewTranscriptionClient()
}
func NewTranscriptionClientWithBaseURI(baseURI string) TranscriptionClient {
	return original.NewTranscriptionClientWithBaseURI(baseURI)
}
func NewTranscriptionsClient() TranscriptionsClient {
	return original.NewTranscriptionsClient()
}
func NewTranscriptionsClientWithBaseURI(baseURI string) TranscriptionsClient {
	return original.NewTranscriptionsClientWithBaseURI(baseURI)
}
func NewVoiceDatasetClient() VoiceDatasetClient {
	return original.NewVoiceDatasetClient()
}
func NewVoiceDatasetClientWithBaseURI(baseURI string) VoiceDatasetClient {
	return original.NewVoiceDatasetClientWithBaseURI(baseURI)
}
func NewVoiceDatasetsClient() VoiceDatasetsClient {
	return original.NewVoiceDatasetsClient()
}
func NewVoiceDatasetsClientWithBaseURI(baseURI string) VoiceDatasetsClient {
	return original.NewVoiceDatasetsClientWithBaseURI(baseURI)
}
func NewVoiceDeploymentClient() VoiceDeploymentClient {
	return original.NewVoiceDeploymentClient()
}
func NewVoiceDeploymentClientWithBaseURI(baseURI string) VoiceDeploymentClient {
	return original.NewVoiceDeploymentClientWithBaseURI(baseURI)
}
func NewVoiceDeploymentsClient() VoiceDeploymentsClient {
	return original.NewVoiceDeploymentsClient()
}
func NewVoiceDeploymentsClientWithBaseURI(baseURI string) VoiceDeploymentsClient {
	return original.NewVoiceDeploymentsClientWithBaseURI(baseURI)
}
func NewVoiceEndpointClient() VoiceEndpointClient {
	return original.NewVoiceEndpointClient()
}
func NewVoiceEndpointClientWithBaseURI(baseURI string) VoiceEndpointClient {
	return original.NewVoiceEndpointClientWithBaseURI(baseURI)
}
func NewVoiceModelClient() VoiceModelClient {
	return original.NewVoiceModelClient()
}
func NewVoiceModelClientWithBaseURI(baseURI string) VoiceModelClient {
	return original.NewVoiceModelClientWithBaseURI(baseURI)
}
func NewVoiceModelsClient() VoiceModelsClient {
	return original.NewVoiceModelsClient()
}
func NewVoiceModelsClientWithBaseURI(baseURI string) VoiceModelsClient {
	return original.NewVoiceModelsClientWithBaseURI(baseURI)
}
func NewVoiceTestClient() VoiceTestClient {
	return original.NewVoiceTestClient()
}
func NewVoiceTestClientWithBaseURI(baseURI string) VoiceTestClient {
	return original.NewVoiceTestClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleDataImportKind1Values() []DataImportKind1 {
	return original.PossibleDataImportKind1Values()
}
func PossibleDataImportKindValues() []DataImportKind {
	return original.PossibleDataImportKindValues()
}
func PossibleEndpointKindValues() []EndpointKind {
	return original.PossibleEndpointKindValues()
}
func PossibleModelKind1Values() []ModelKind1 {
	return original.PossibleModelKind1Values()
}
func PossibleModelKind2Values() []ModelKind2 {
	return original.PossibleModelKind2Values()
}
func PossibleModelKindValues() []ModelKind {
	return original.PossibleModelKindValues()
}
func PossibleStatus1Values() []Status1 {
	return original.PossibleStatus1Values()
}
func PossibleStatus2Values() []Status2 {
	return original.PossibleStatus2Values()
}
func PossibleStatus3Values() []Status3 {
	return original.PossibleStatus3Values()
}
func PossibleStatus4Values() []Status4 {
	return original.PossibleStatus4Values()
}
func PossibleStatus5Values() []Status5 {
	return original.PossibleStatus5Values()
}
func PossibleStatus6Values() []Status6 {
	return original.PossibleStatus6Values()
}
func PossibleStatus7Values() []Status7 {
	return original.PossibleStatus7Values()
}
func PossibleStatus8Values() []Status8 {
	return original.PossibleStatus8Values()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleVoiceTestKind1Values() []VoiceTestKind1 {
	return original.PossibleVoiceTestKind1Values()
}
func PossibleVoiceTestKindValues() []VoiceTestKind {
	return original.PossibleVoiceTestKindValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
