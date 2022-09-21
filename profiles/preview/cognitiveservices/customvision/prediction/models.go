//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package prediction

import original "github.com/Azure/dev/azure-sdk-for-go/services/cognitiveservices/v3.1/customvision/prediction"

type CustomVisionErrorCodes = original.CustomVisionErrorCodes

const (
	BadRequest                                                  CustomVisionErrorCodes = original.BadRequest
	BadRequestCannotMigrateProjectWithName                      CustomVisionErrorCodes = original.BadRequestCannotMigrateProjectWithName
	BadRequestClassificationTrainingValidationFailed            CustomVisionErrorCodes = original.BadRequestClassificationTrainingValidationFailed
	BadRequestCustomerManagedKeyRevoked                         CustomVisionErrorCodes = original.BadRequestCustomerManagedKeyRevoked
	BadRequestDetectionTrainingNotAllowNegativeTag              CustomVisionErrorCodes = original.BadRequestDetectionTrainingNotAllowNegativeTag
	BadRequestDetectionTrainingValidationFailed                 CustomVisionErrorCodes = original.BadRequestDetectionTrainingValidationFailed
	BadRequestDomainNotSupportedForAdvancedTraining             CustomVisionErrorCodes = original.BadRequestDomainNotSupportedForAdvancedTraining
	BadRequestExceededBatchSize                                 CustomVisionErrorCodes = original.BadRequestExceededBatchSize
	BadRequestExceededQuota                                     CustomVisionErrorCodes = original.BadRequestExceededQuota
	BadRequestExceedIterationPerProjectLimit                    CustomVisionErrorCodes = original.BadRequestExceedIterationPerProjectLimit
	BadRequestExceedProjectLimit                                CustomVisionErrorCodes = original.BadRequestExceedProjectLimit
	BadRequestExceedTagPerImageLimit                            CustomVisionErrorCodes = original.BadRequestExceedTagPerImageLimit
	BadRequestExceedTagPerProjectLimit                          CustomVisionErrorCodes = original.BadRequestExceedTagPerProjectLimit
	BadRequestExportAlreadyInProgress                           CustomVisionErrorCodes = original.BadRequestExportAlreadyInProgress
	BadRequestExportPlatformNotSupportedForAdvancedTraining     CustomVisionErrorCodes = original.BadRequestExportPlatformNotSupportedForAdvancedTraining
	BadRequestExportValidationFailed                            CustomVisionErrorCodes = original.BadRequestExportValidationFailed
	BadRequestExportWhileTraining                               CustomVisionErrorCodes = original.BadRequestExportWhileTraining
	BadRequestImageBatch                                        CustomVisionErrorCodes = original.BadRequestImageBatch
	BadRequestImageDimensions                                   CustomVisionErrorCodes = original.BadRequestImageDimensions
	BadRequestImageExceededCount                                CustomVisionErrorCodes = original.BadRequestImageExceededCount
	BadRequestImageFormat                                       CustomVisionErrorCodes = original.BadRequestImageFormat
	BadRequestImageMetadataKey                                  CustomVisionErrorCodes = original.BadRequestImageMetadataKey
	BadRequestImageMetadataValue                                CustomVisionErrorCodes = original.BadRequestImageMetadataValue
	BadRequestImageRegions                                      CustomVisionErrorCodes = original.BadRequestImageRegions
	BadRequestImageSizeBytes                                    CustomVisionErrorCodes = original.BadRequestImageSizeBytes
	BadRequestImageStream                                       CustomVisionErrorCodes = original.BadRequestImageStream
	BadRequestImageTags                                         CustomVisionErrorCodes = original.BadRequestImageTags
	BadRequestImageURL                                          CustomVisionErrorCodes = original.BadRequestImageURL
	BadRequestInvalid                                           CustomVisionErrorCodes = original.BadRequestInvalid
	BadRequestInvalidArtifactURI                                CustomVisionErrorCodes = original.BadRequestInvalidArtifactURI
	BadRequestInvalidEmailAddress                               CustomVisionErrorCodes = original.BadRequestInvalidEmailAddress
	BadRequestInvalidIds                                        CustomVisionErrorCodes = original.BadRequestInvalidIds
	BadRequestInvalidImportToken                                CustomVisionErrorCodes = original.BadRequestInvalidImportToken
	BadRequestInvalidPublishName                                CustomVisionErrorCodes = original.BadRequestInvalidPublishName
	BadRequestInvalidPublishTarget                              CustomVisionErrorCodes = original.BadRequestInvalidPublishTarget
	BadRequestIterationDescription                              CustomVisionErrorCodes = original.BadRequestIterationDescription
	BadRequestIterationIsNotTrained                             CustomVisionErrorCodes = original.BadRequestIterationIsNotTrained
	BadRequestIterationIsPublished                              CustomVisionErrorCodes = original.BadRequestIterationIsPublished
	BadRequestIterationName                                     CustomVisionErrorCodes = original.BadRequestIterationName
	BadRequestIterationNameNotUnique                            CustomVisionErrorCodes = original.BadRequestIterationNameNotUnique
	BadRequestIterationNotPublished                             CustomVisionErrorCodes = original.BadRequestIterationNotPublished
	BadRequestIterationValidationFailed                         CustomVisionErrorCodes = original.BadRequestIterationValidationFailed
	BadRequestMultiClassClassificationTrainingValidationFailed  CustomVisionErrorCodes = original.BadRequestMultiClassClassificationTrainingValidationFailed
	BadRequestMultiLabelClassificationTrainingValidationFailed  CustomVisionErrorCodes = original.BadRequestMultiLabelClassificationTrainingValidationFailed
	BadRequestMultipleGeneralProductTag                         CustomVisionErrorCodes = original.BadRequestMultipleGeneralProductTag
	BadRequestMultipleNegativeTag                               CustomVisionErrorCodes = original.BadRequestMultipleNegativeTag
	BadRequestNegativeAndRegularTagOnSameImage                  CustomVisionErrorCodes = original.BadRequestNegativeAndRegularTagOnSameImage
	BadRequestNotLimitedTrial                                   CustomVisionErrorCodes = original.BadRequestNotLimitedTrial
	BadRequestNotSupported                                      CustomVisionErrorCodes = original.BadRequestNotSupported
	BadRequestOperationNotSupported                             CustomVisionErrorCodes = original.BadRequestOperationNotSupported
	BadRequestPredictionIdsExceededCount                        CustomVisionErrorCodes = original.BadRequestPredictionIdsExceededCount
	BadRequestPredictionIdsMissing                              CustomVisionErrorCodes = original.BadRequestPredictionIdsMissing
	BadRequestPredictionInvalidApplicationName                  CustomVisionErrorCodes = original.BadRequestPredictionInvalidApplicationName
	BadRequestPredictionInvalidQueryParameters                  CustomVisionErrorCodes = original.BadRequestPredictionInvalidQueryParameters
	BadRequestPredictionResultsExceededCount                    CustomVisionErrorCodes = original.BadRequestPredictionResultsExceededCount
	BadRequestPredictionTagsExceededCount                       CustomVisionErrorCodes = original.BadRequestPredictionTagsExceededCount
	BadRequestProjectDescription                                CustomVisionErrorCodes = original.BadRequestProjectDescription
	BadRequestProjectDuplicated                                 CustomVisionErrorCodes = original.BadRequestProjectDuplicated
	BadRequestProjectImagePreprocessingSettings                 CustomVisionErrorCodes = original.BadRequestProjectImagePreprocessingSettings
	BadRequestProjectName                                       CustomVisionErrorCodes = original.BadRequestProjectName
	BadRequestProjectNameNotUnique                              CustomVisionErrorCodes = original.BadRequestProjectNameNotUnique
	BadRequestProjectUnknownClassification                      CustomVisionErrorCodes = original.BadRequestProjectUnknownClassification
	BadRequestProjectUnknownDomain                              CustomVisionErrorCodes = original.BadRequestProjectUnknownDomain
	BadRequestProjectUnsupportedDomainTypeChange                CustomVisionErrorCodes = original.BadRequestProjectUnsupportedDomainTypeChange
	BadRequestProjectUnsupportedExportPlatform                  CustomVisionErrorCodes = original.BadRequestProjectUnsupportedExportPlatform
	BadRequestRequiredParamIsNull                               CustomVisionErrorCodes = original.BadRequestRequiredParamIsNull
	BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining CustomVisionErrorCodes = original.BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining
	BadRequestSubscriptionAPI                                   CustomVisionErrorCodes = original.BadRequestSubscriptionAPI
	BadRequestTagDescription                                    CustomVisionErrorCodes = original.BadRequestTagDescription
	BadRequestTagName                                           CustomVisionErrorCodes = original.BadRequestTagName
	BadRequestTagNameNotUnique                                  CustomVisionErrorCodes = original.BadRequestTagNameNotUnique
	BadRequestTagType                                           CustomVisionErrorCodes = original.BadRequestTagType
	BadRequestTrainingAlreadyInProgress                         CustomVisionErrorCodes = original.BadRequestTrainingAlreadyInProgress
	BadRequestTrainingNotNeeded                                 CustomVisionErrorCodes = original.BadRequestTrainingNotNeeded
	BadRequestTrainingNotNeededButTrainingPipelineUpdated       CustomVisionErrorCodes = original.BadRequestTrainingNotNeededButTrainingPipelineUpdated
	BadRequestTrainingValidationFailed                          CustomVisionErrorCodes = original.BadRequestTrainingValidationFailed
	BadRequestUnpublishFailed                                   CustomVisionErrorCodes = original.BadRequestUnpublishFailed
	BadRequestUnsupportedDomain                                 CustomVisionErrorCodes = original.BadRequestUnsupportedDomain
	BadRequestWorkspaceCannotBeModified                         CustomVisionErrorCodes = original.BadRequestWorkspaceCannotBeModified
	BadRequestWorkspaceNotDeletable                             CustomVisionErrorCodes = original.BadRequestWorkspaceNotDeletable
	Conflict                                                    CustomVisionErrorCodes = original.Conflict
	ConflictInvalid                                             CustomVisionErrorCodes = original.ConflictInvalid
	ErrorExporterInvalidClassifier                              CustomVisionErrorCodes = original.ErrorExporterInvalidClassifier
	ErrorExporterInvalidFeaturizer                              CustomVisionErrorCodes = original.ErrorExporterInvalidFeaturizer
	ErrorExporterInvalidPlatform                                CustomVisionErrorCodes = original.ErrorExporterInvalidPlatform
	ErrorFeaturizationAugmentationError                         CustomVisionErrorCodes = original.ErrorFeaturizationAugmentationError
	ErrorFeaturizationAugmentationUnavailable                   CustomVisionErrorCodes = original.ErrorFeaturizationAugmentationUnavailable
	ErrorFeaturizationInvalidFeaturizer                         CustomVisionErrorCodes = original.ErrorFeaturizationInvalidFeaturizer
	ErrorFeaturizationQueueTimeout                              CustomVisionErrorCodes = original.ErrorFeaturizationQueueTimeout
	ErrorFeaturizationServiceUnavailable                        CustomVisionErrorCodes = original.ErrorFeaturizationServiceUnavailable
	ErrorFeaturizationUnrecognizedJob                           CustomVisionErrorCodes = original.ErrorFeaturizationUnrecognizedJob
	ErrorInvalid                                                CustomVisionErrorCodes = original.ErrorInvalid
	ErrorIterationCopyFailed                                    CustomVisionErrorCodes = original.ErrorIterationCopyFailed
	ErrorPrediction                                             CustomVisionErrorCodes = original.ErrorPrediction
	ErrorPredictionModelNotCached                               CustomVisionErrorCodes = original.ErrorPredictionModelNotCached
	ErrorPredictionModelNotFound                                CustomVisionErrorCodes = original.ErrorPredictionModelNotFound
	ErrorPredictionServiceUnavailable                           CustomVisionErrorCodes = original.ErrorPredictionServiceUnavailable
	ErrorPredictionStorage                                      CustomVisionErrorCodes = original.ErrorPredictionStorage
	ErrorPreparePerformanceMigrationFailed                      CustomVisionErrorCodes = original.ErrorPreparePerformanceMigrationFailed
	ErrorProjectExportRequestFailed                             CustomVisionErrorCodes = original.ErrorProjectExportRequestFailed
	ErrorProjectImportRequestFailed                             CustomVisionErrorCodes = original.ErrorProjectImportRequestFailed
	ErrorProjectInvalidDomain                                   CustomVisionErrorCodes = original.ErrorProjectInvalidDomain
	ErrorProjectInvalidPipelineConfiguration                    CustomVisionErrorCodes = original.ErrorProjectInvalidPipelineConfiguration
	ErrorProjectInvalidWorkspace                                CustomVisionErrorCodes = original.ErrorProjectInvalidWorkspace
	ErrorProjectTrainingRequestFailed                           CustomVisionErrorCodes = original.ErrorProjectTrainingRequestFailed
	ErrorRegionProposal                                         CustomVisionErrorCodes = original.ErrorRegionProposal
	ErrorUnknown                                                CustomVisionErrorCodes = original.ErrorUnknown
	ErrorUnknownBaseModel                                       CustomVisionErrorCodes = original.ErrorUnknownBaseModel
	Forbidden                                                   CustomVisionErrorCodes = original.Forbidden
	ForbiddenDRModeEnabled                                      CustomVisionErrorCodes = original.ForbiddenDRModeEnabled
	ForbiddenInvalid                                            CustomVisionErrorCodes = original.ForbiddenInvalid
	ForbiddenUser                                               CustomVisionErrorCodes = original.ForbiddenUser
	ForbiddenUserDisabled                                       CustomVisionErrorCodes = original.ForbiddenUserDisabled
	ForbiddenUserDoesNotExist                                   CustomVisionErrorCodes = original.ForbiddenUserDoesNotExist
	ForbiddenUserInsufficientCapability                         CustomVisionErrorCodes = original.ForbiddenUserInsufficientCapability
	ForbiddenUserResource                                       CustomVisionErrorCodes = original.ForbiddenUserResource
	ForbiddenUserSignupAllowanceExceeded                        CustomVisionErrorCodes = original.ForbiddenUserSignupAllowanceExceeded
	ForbiddenUserSignupDisabled                                 CustomVisionErrorCodes = original.ForbiddenUserSignupDisabled
	NoError                                                     CustomVisionErrorCodes = original.NoError
	NotFound                                                    CustomVisionErrorCodes = original.NotFound
	NotFoundApimSubscription                                    CustomVisionErrorCodes = original.NotFoundApimSubscription
	NotFoundDomain                                              CustomVisionErrorCodes = original.NotFoundDomain
	NotFoundImage                                               CustomVisionErrorCodes = original.NotFoundImage
	NotFoundInvalid                                             CustomVisionErrorCodes = original.NotFoundInvalid
	NotFoundIteration                                           CustomVisionErrorCodes = original.NotFoundIteration
	NotFoundIterationPerformance                                CustomVisionErrorCodes = original.NotFoundIterationPerformance
	NotFoundProject                                             CustomVisionErrorCodes = original.NotFoundProject
	NotFoundProjectDefaultIteration                             CustomVisionErrorCodes = original.NotFoundProjectDefaultIteration
	NotFoundTag                                                 CustomVisionErrorCodes = original.NotFoundTag
	UnsupportedMediaType                                        CustomVisionErrorCodes = original.UnsupportedMediaType
)

type TagType = original.TagType

const (
	GeneralProduct TagType = original.GeneralProduct
	Negative       TagType = original.Negative
	Regular        TagType = original.Regular
)

type BaseClient = original.BaseClient
type BoundingBox = original.BoundingBox
type CustomVisionError = original.CustomVisionError
type ImagePrediction = original.ImagePrediction
type ImageURL = original.ImageURL
type Model = original.Model

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleCustomVisionErrorCodesValues() []CustomVisionErrorCodes {
	return original.PossibleCustomVisionErrorCodesValues()
}
func PossibleTagTypeValues() []TagType {
	return original.PossibleTagTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
