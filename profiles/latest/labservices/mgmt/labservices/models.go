//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package labservices

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/labservices/mgmt/2018-10-15/labservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AddRemove = original.AddRemove

const (
	Add    AddRemove = original.Add
	Remove AddRemove = original.Remove
)

type ConfigurationState = original.ConfigurationState

const (
	Completed     ConfigurationState = original.Completed
	NotApplicable ConfigurationState = original.NotApplicable
)

type LabUserAccessMode = original.LabUserAccessMode

const (
	Open       LabUserAccessMode = original.Open
	Restricted LabUserAccessMode = original.Restricted
)

type ManagedLabVMSize = original.ManagedLabVMSize

const (
	Basic       ManagedLabVMSize = original.Basic
	Performance ManagedLabVMSize = original.Performance
	Standard    ManagedLabVMSize = original.Standard
)

type PublishingState = original.PublishingState

const (
	Draft         PublishingState = original.Draft
	Published     PublishingState = original.Published
	PublishFailed PublishingState = original.PublishFailed
	Publishing    PublishingState = original.Publishing
	Scaling       PublishingState = original.Scaling
)

type AddUsersPayload = original.AddUsersPayload
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CreateLabProperties = original.CreateLabProperties
type Environment = original.Environment
type EnvironmentDetails = original.EnvironmentDetails
type EnvironmentFragment = original.EnvironmentFragment
type EnvironmentOperationsPayload = original.EnvironmentOperationsPayload
type EnvironmentProperties = original.EnvironmentProperties
type EnvironmentPropertiesFragment = original.EnvironmentPropertiesFragment
type EnvironmentSetting = original.EnvironmentSetting
type EnvironmentSettingCreationParameters = original.EnvironmentSettingCreationParameters
type EnvironmentSettingFragment = original.EnvironmentSettingFragment
type EnvironmentSettingProperties = original.EnvironmentSettingProperties
type EnvironmentSettingPropertiesFragment = original.EnvironmentSettingPropertiesFragment
type EnvironmentSettingsClient = original.EnvironmentSettingsClient
type EnvironmentSettingsCreateOrUpdateFuture = original.EnvironmentSettingsCreateOrUpdateFuture
type EnvironmentSettingsDeleteFuture = original.EnvironmentSettingsDeleteFuture
type EnvironmentSettingsStartFuture = original.EnvironmentSettingsStartFuture
type EnvironmentSettingsStopFuture = original.EnvironmentSettingsStopFuture
type EnvironmentSize = original.EnvironmentSize
type EnvironmentSizeFragment = original.EnvironmentSizeFragment
type EnvironmentsClient = original.EnvironmentsClient
type EnvironmentsDeleteFuture = original.EnvironmentsDeleteFuture
type EnvironmentsResetPasswordFuture = original.EnvironmentsResetPasswordFuture
type EnvironmentsStartFuture = original.EnvironmentsStartFuture
type EnvironmentsStopFuture = original.EnvironmentsStopFuture
type GalleryImage = original.GalleryImage
type GalleryImageFragment = original.GalleryImageFragment
type GalleryImageProperties = original.GalleryImageProperties
type GalleryImagePropertiesFragment = original.GalleryImagePropertiesFragment
type GalleryImageReference = original.GalleryImageReference
type GalleryImageReferenceFragment = original.GalleryImageReferenceFragment
type GalleryImagesClient = original.GalleryImagesClient
type GetEnvironmentResponse = original.GetEnvironmentResponse
type GetPersonalPreferencesResponse = original.GetPersonalPreferencesResponse
type GetRegionalAvailabilityResponse = original.GetRegionalAvailabilityResponse
type GlobalUsersClient = original.GlobalUsersClient
type GlobalUsersResetPasswordFuture = original.GlobalUsersResetPasswordFuture
type GlobalUsersStartEnvironmentFuture = original.GlobalUsersStartEnvironmentFuture
type GlobalUsersStopEnvironmentFuture = original.GlobalUsersStopEnvironmentFuture
type Lab = original.Lab
type LabAccount = original.LabAccount
type LabAccountFragment = original.LabAccountFragment
type LabAccountProperties = original.LabAccountProperties
type LabAccountPropertiesFragment = original.LabAccountPropertiesFragment
type LabAccountsClient = original.LabAccountsClient
type LabAccountsDeleteFuture = original.LabAccountsDeleteFuture
type LabCreationParameters = original.LabCreationParameters
type LabDetails = original.LabDetails
type LabFragment = original.LabFragment
type LabProperties = original.LabProperties
type LabPropertiesFragment = original.LabPropertiesFragment
type LabsClient = original.LabsClient
type LabsDeleteFuture = original.LabsDeleteFuture
type LatestOperationResult = original.LatestOperationResult
type ListEnvironmentsPayload = original.ListEnvironmentsPayload
type ListEnvironmentsResponse = original.ListEnvironmentsResponse
type ListLabsResponse = original.ListLabsResponse
type NetworkInterface = original.NetworkInterface
type OperationBatchStatusPayload = original.OperationBatchStatusPayload
type OperationBatchStatusResponse = original.OperationBatchStatusResponse
type OperationBatchStatusResponseItem = original.OperationBatchStatusResponseItem
type OperationError = original.OperationError
type OperationMetadata = original.OperationMetadata
type OperationMetadataDisplay = original.OperationMetadataDisplay
type OperationResult = original.OperationResult
type OperationStatusPayload = original.OperationStatusPayload
type OperationStatusResponse = original.OperationStatusResponse
type OperationsClient = original.OperationsClient
type PersonalPreferencesOperationsPayload = original.PersonalPreferencesOperationsPayload
type ProviderOperationResult = original.ProviderOperationResult
type ProviderOperationResultIterator = original.ProviderOperationResultIterator
type ProviderOperationResultPage = original.ProviderOperationResultPage
type ProviderOperationsClient = original.ProviderOperationsClient
type PublishPayload = original.PublishPayload
type ReferenceVM = original.ReferenceVM
type ReferenceVMCreationParameters = original.ReferenceVMCreationParameters
type ReferenceVMFragment = original.ReferenceVMFragment
type RegionalAvailability = original.RegionalAvailability
type RegisterPayload = original.RegisterPayload
type ResetPasswordPayload = original.ResetPasswordPayload
type Resource = original.Resource
type ResourceSet = original.ResourceSet
type ResourceSetFragment = original.ResourceSetFragment
type ResourceSettingCreationParameters = original.ResourceSettingCreationParameters
type ResourceSettings = original.ResourceSettings
type ResourceSettingsFragment = original.ResourceSettingsFragment
type ResponseWithContinuationEnvironment = original.ResponseWithContinuationEnvironment
type ResponseWithContinuationEnvironmentIterator = original.ResponseWithContinuationEnvironmentIterator
type ResponseWithContinuationEnvironmentPage = original.ResponseWithContinuationEnvironmentPage
type ResponseWithContinuationEnvironmentSetting = original.ResponseWithContinuationEnvironmentSetting
type ResponseWithContinuationEnvironmentSettingIterator = original.ResponseWithContinuationEnvironmentSettingIterator
type ResponseWithContinuationEnvironmentSettingPage = original.ResponseWithContinuationEnvironmentSettingPage
type ResponseWithContinuationGalleryImage = original.ResponseWithContinuationGalleryImage
type ResponseWithContinuationGalleryImageIterator = original.ResponseWithContinuationGalleryImageIterator
type ResponseWithContinuationGalleryImagePage = original.ResponseWithContinuationGalleryImagePage
type ResponseWithContinuationLab = original.ResponseWithContinuationLab
type ResponseWithContinuationLabAccount = original.ResponseWithContinuationLabAccount
type ResponseWithContinuationLabAccountIterator = original.ResponseWithContinuationLabAccountIterator
type ResponseWithContinuationLabAccountPage = original.ResponseWithContinuationLabAccountPage
type ResponseWithContinuationLabIterator = original.ResponseWithContinuationLabIterator
type ResponseWithContinuationLabPage = original.ResponseWithContinuationLabPage
type ResponseWithContinuationUser = original.ResponseWithContinuationUser
type ResponseWithContinuationUserIterator = original.ResponseWithContinuationUserIterator
type ResponseWithContinuationUserPage = original.ResponseWithContinuationUserPage
type SizeAvailability = original.SizeAvailability
type SizeConfigurationProperties = original.SizeConfigurationProperties
type SizeConfigurationPropertiesFragment = original.SizeConfigurationPropertiesFragment
type SizeInfo = original.SizeInfo
type SizeInfoFragment = original.SizeInfoFragment
type User = original.User
type UserFragment = original.UserFragment
type UserProperties = original.UserProperties
type UserPropertiesFragment = original.UserPropertiesFragment
type UsersClient = original.UsersClient
type UsersDeleteFuture = original.UsersDeleteFuture
type VMStateDetails = original.VMStateDetails
type VirtualMachineDetails = original.VirtualMachineDetails

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEnvironmentSettingsClient(subscriptionID string) EnvironmentSettingsClient {
	return original.NewEnvironmentSettingsClient(subscriptionID)
}
func NewEnvironmentSettingsClientWithBaseURI(baseURI string, subscriptionID string) EnvironmentSettingsClient {
	return original.NewEnvironmentSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEnvironmentsClient(subscriptionID string) EnvironmentsClient {
	return original.NewEnvironmentsClient(subscriptionID)
}
func NewEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) EnvironmentsClient {
	return original.NewEnvironmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGalleryImagesClient(subscriptionID string) GalleryImagesClient {
	return original.NewGalleryImagesClient(subscriptionID)
}
func NewGalleryImagesClientWithBaseURI(baseURI string, subscriptionID string) GalleryImagesClient {
	return original.NewGalleryImagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewGlobalUsersClient(subscriptionID string) GlobalUsersClient {
	return original.NewGlobalUsersClient(subscriptionID)
}
func NewGlobalUsersClientWithBaseURI(baseURI string, subscriptionID string) GlobalUsersClient {
	return original.NewGlobalUsersClientWithBaseURI(baseURI, subscriptionID)
}
func NewLabAccountsClient(subscriptionID string) LabAccountsClient {
	return original.NewLabAccountsClient(subscriptionID)
}
func NewLabAccountsClientWithBaseURI(baseURI string, subscriptionID string) LabAccountsClient {
	return original.NewLabAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLabsClient(subscriptionID string) LabsClient {
	return original.NewLabsClient(subscriptionID)
}
func NewLabsClientWithBaseURI(baseURI string, subscriptionID string) LabsClient {
	return original.NewLabsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderOperationResultIterator(page ProviderOperationResultPage) ProviderOperationResultIterator {
	return original.NewProviderOperationResultIterator(page)
}
func NewProviderOperationResultPage(cur ProviderOperationResult, getNextPage func(context.Context, ProviderOperationResult) (ProviderOperationResult, error)) ProviderOperationResultPage {
	return original.NewProviderOperationResultPage(cur, getNextPage)
}
func NewProviderOperationsClient(subscriptionID string) ProviderOperationsClient {
	return original.NewProviderOperationsClient(subscriptionID)
}
func NewProviderOperationsClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationsClient {
	return original.NewProviderOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResponseWithContinuationEnvironmentIterator(page ResponseWithContinuationEnvironmentPage) ResponseWithContinuationEnvironmentIterator {
	return original.NewResponseWithContinuationEnvironmentIterator(page)
}
func NewResponseWithContinuationEnvironmentPage(cur ResponseWithContinuationEnvironment, getNextPage func(context.Context, ResponseWithContinuationEnvironment) (ResponseWithContinuationEnvironment, error)) ResponseWithContinuationEnvironmentPage {
	return original.NewResponseWithContinuationEnvironmentPage(cur, getNextPage)
}
func NewResponseWithContinuationEnvironmentSettingIterator(page ResponseWithContinuationEnvironmentSettingPage) ResponseWithContinuationEnvironmentSettingIterator {
	return original.NewResponseWithContinuationEnvironmentSettingIterator(page)
}
func NewResponseWithContinuationEnvironmentSettingPage(cur ResponseWithContinuationEnvironmentSetting, getNextPage func(context.Context, ResponseWithContinuationEnvironmentSetting) (ResponseWithContinuationEnvironmentSetting, error)) ResponseWithContinuationEnvironmentSettingPage {
	return original.NewResponseWithContinuationEnvironmentSettingPage(cur, getNextPage)
}
func NewResponseWithContinuationGalleryImageIterator(page ResponseWithContinuationGalleryImagePage) ResponseWithContinuationGalleryImageIterator {
	return original.NewResponseWithContinuationGalleryImageIterator(page)
}
func NewResponseWithContinuationGalleryImagePage(cur ResponseWithContinuationGalleryImage, getNextPage func(context.Context, ResponseWithContinuationGalleryImage) (ResponseWithContinuationGalleryImage, error)) ResponseWithContinuationGalleryImagePage {
	return original.NewResponseWithContinuationGalleryImagePage(cur, getNextPage)
}
func NewResponseWithContinuationLabAccountIterator(page ResponseWithContinuationLabAccountPage) ResponseWithContinuationLabAccountIterator {
	return original.NewResponseWithContinuationLabAccountIterator(page)
}
func NewResponseWithContinuationLabAccountPage(cur ResponseWithContinuationLabAccount, getNextPage func(context.Context, ResponseWithContinuationLabAccount) (ResponseWithContinuationLabAccount, error)) ResponseWithContinuationLabAccountPage {
	return original.NewResponseWithContinuationLabAccountPage(cur, getNextPage)
}
func NewResponseWithContinuationLabIterator(page ResponseWithContinuationLabPage) ResponseWithContinuationLabIterator {
	return original.NewResponseWithContinuationLabIterator(page)
}
func NewResponseWithContinuationLabPage(cur ResponseWithContinuationLab, getNextPage func(context.Context, ResponseWithContinuationLab) (ResponseWithContinuationLab, error)) ResponseWithContinuationLabPage {
	return original.NewResponseWithContinuationLabPage(cur, getNextPage)
}
func NewResponseWithContinuationUserIterator(page ResponseWithContinuationUserPage) ResponseWithContinuationUserIterator {
	return original.NewResponseWithContinuationUserIterator(page)
}
func NewResponseWithContinuationUserPage(cur ResponseWithContinuationUser, getNextPage func(context.Context, ResponseWithContinuationUser) (ResponseWithContinuationUser, error)) ResponseWithContinuationUserPage {
	return original.NewResponseWithContinuationUserPage(cur, getNextPage)
}
func NewUsersClient(subscriptionID string) UsersClient {
	return original.NewUsersClient(subscriptionID)
}
func NewUsersClientWithBaseURI(baseURI string, subscriptionID string) UsersClient {
	return original.NewUsersClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAddRemoveValues() []AddRemove {
	return original.PossibleAddRemoveValues()
}
func PossibleConfigurationStateValues() []ConfigurationState {
	return original.PossibleConfigurationStateValues()
}
func PossibleLabUserAccessModeValues() []LabUserAccessMode {
	return original.PossibleLabUserAccessModeValues()
}
func PossibleManagedLabVMSizeValues() []ManagedLabVMSize {
	return original.PossibleManagedLabVMSizeValues()
}
func PossiblePublishingStateValues() []PublishingState {
	return original.PossiblePublishingStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
