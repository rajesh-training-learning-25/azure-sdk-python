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

package contentmoderator

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/contentmoderator"

type BaseClient = original.BaseClient

func New(baseURL AzureRegionBaseURL) BaseClient {
	return original.New(baseURL)
}
func NewWithoutDefaults(baseURL AzureRegionBaseURL) BaseClient {
	return original.NewWithoutDefaults(baseURL)
}

type ImageModerationClient = original.ImageModerationClient

func NewImageModerationClient(baseURL AzureRegionBaseURL) ImageModerationClient {
	return original.NewImageModerationClient(baseURL)
}

type ListManagementImageClient = original.ListManagementImageClient

func NewListManagementImageClient(baseURL AzureRegionBaseURL) ListManagementImageClient {
	return original.NewListManagementImageClient(baseURL)
}

type ListManagementImageListsClient = original.ListManagementImageListsClient

func NewListManagementImageListsClient(baseURL AzureRegionBaseURL) ListManagementImageListsClient {
	return original.NewListManagementImageListsClient(baseURL)
}

type ListManagementTermClient = original.ListManagementTermClient

func NewListManagementTermClient(baseURL AzureRegionBaseURL) ListManagementTermClient {
	return original.NewListManagementTermClient(baseURL)
}

type ListManagementTermListsClient = original.ListManagementTermListsClient

func NewListManagementTermListsClient(baseURL AzureRegionBaseURL) ListManagementTermListsClient {
	return original.NewListManagementTermListsClient(baseURL)
}

type AzureRegionBaseURL = original.AzureRegionBaseURL

const (
	Australiaeastapicognitivemicrosoftcom  AzureRegionBaseURL = original.Australiaeastapicognitivemicrosoftcom
	Brazilsouthapicognitivemicrosoftcom    AzureRegionBaseURL = original.Brazilsouthapicognitivemicrosoftcom
	ContentmoderatortestazureApinet        AzureRegionBaseURL = original.ContentmoderatortestazureApinet
	Eastasiaapicognitivemicrosoftcom       AzureRegionBaseURL = original.Eastasiaapicognitivemicrosoftcom
	Eastus2apicognitivemicrosoftcom        AzureRegionBaseURL = original.Eastus2apicognitivemicrosoftcom
	Eastusapicognitivemicrosoftcom         AzureRegionBaseURL = original.Eastusapicognitivemicrosoftcom
	Northeuropeapicognitivemicrosoftcom    AzureRegionBaseURL = original.Northeuropeapicognitivemicrosoftcom
	Southcentralusapicognitivemicrosoftcom AzureRegionBaseURL = original.Southcentralusapicognitivemicrosoftcom
	Southeastasiaapicognitivemicrosoftcom  AzureRegionBaseURL = original.Southeastasiaapicognitivemicrosoftcom
	Westcentralusapicognitivemicrosoftcom  AzureRegionBaseURL = original.Westcentralusapicognitivemicrosoftcom
	Westeuropeapicognitivemicrosoftcom     AzureRegionBaseURL = original.Westeuropeapicognitivemicrosoftcom
	Westus2apicognitivemicrosoftcom        AzureRegionBaseURL = original.Westus2apicognitivemicrosoftcom
	Westusapicognitivemicrosoftcom         AzureRegionBaseURL = original.Westusapicognitivemicrosoftcom
)

func PossibleAzureRegionBaseURLValues() []AzureRegionBaseURL {
	return original.PossibleAzureRegionBaseURLValues()
}

type StatusEnum = original.StatusEnum

const (
	Complete    StatusEnum = original.Complete
	Pending     StatusEnum = original.Pending
	Unpublished StatusEnum = original.Unpublished
)

func PossibleStatusEnumValues() []StatusEnum {
	return original.PossibleStatusEnumValues()
}

type Type = original.Type

const (
	TypeImage Type = original.TypeImage
	TypeText  Type = original.TypeText
)

func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}

type Address = original.Address
type APIError = original.APIError
type Body = original.Body
type BodyMetadata = original.BodyMetadata
type BodyModel = original.BodyModel
type Candidate = original.Candidate
type Classification = original.Classification
type ClassificationCategory1 = original.ClassificationCategory1
type ClassificationCategory2 = original.ClassificationCategory2
type ClassificationCategory3 = original.ClassificationCategory3
type Content = original.Content
type CreateReviewBodyItem = original.CreateReviewBodyItem
type CreateReviewBodyItemMetadataItem = original.CreateReviewBodyItemMetadataItem
type CreateVideoReviewsBodyItem = original.CreateVideoReviewsBodyItem
type CreateVideoReviewsBodyItemMetadataItem = original.CreateVideoReviewsBodyItemMetadataItem
type CreateVideoReviewsBodyItemVideoFramesItem = original.CreateVideoReviewsBodyItemVideoFramesItem
type CreateVideoReviewsBodyItemVideoFramesItemMetadataItem = original.CreateVideoReviewsBodyItemVideoFramesItemMetadataItem
type CreateVideoReviewsBodyItemVideoFramesItemReviewerResultTagsItem = original.CreateVideoReviewsBodyItemVideoFramesItemReviewerResultTagsItem
type DetectedLanguage = original.DetectedLanguage
type DetectedTerms = original.DetectedTerms
type Email = original.Email
type Error = original.Error
type Evaluate = original.Evaluate
type Face = original.Face
type FoundFaces = original.FoundFaces
type Frame = original.Frame
type Frames = original.Frames
type Image = original.Image
type ImageAdditionalInfoItem = original.ImageAdditionalInfoItem
type ImageIds = original.ImageIds
type ImageList = original.ImageList
type ImageListMetadata = original.ImageListMetadata
type IPA = original.IPA
type Job = original.Job
type JobExecutionReportDetails = original.JobExecutionReportDetails
type JobID = original.JobID
type JobListResult = original.JobListResult
type KeyValuePair = original.KeyValuePair
type ListImageList = original.ListImageList
type ListString = original.ListString
type ListTermList = original.ListTermList
type Match = original.Match
type MatchResponse = original.MatchResponse
type OCR = original.OCR
type Phone = original.Phone
type PII = original.PII
type RefreshIndex = original.RefreshIndex
type RefreshIndexAdvancedInfoItem = original.RefreshIndexAdvancedInfoItem
type Review = original.Review
type Screen = original.Screen
type SetObject = original.SetObject
type Status = original.Status
type String = original.String
type Tag = original.Tag
type TermList = original.TermList
type TermListMetadata = original.TermListMetadata
type Terms = original.Terms
type TermsData = original.TermsData
type TermsInList = original.TermsInList
type TermsPaging = original.TermsPaging
type TranscriptModerationBodyItem = original.TranscriptModerationBodyItem
type TranscriptModerationBodyItemTermsItem = original.TranscriptModerationBodyItemTermsItem
type VideoFrameBodyItem = original.VideoFrameBodyItem
type VideoFrameBodyItemMetadataItem = original.VideoFrameBodyItemMetadataItem
type VideoFrameBodyItemReviewerResultTagsItem = original.VideoFrameBodyItemReviewerResultTagsItem
type ReviewsClient = original.ReviewsClient

func NewReviewsClient(baseURL AzureRegionBaseURL) ReviewsClient {
	return original.NewReviewsClient(baseURL)
}

type TextModerationClient = original.TextModerationClient

func NewTextModerationClient(baseURL AzureRegionBaseURL) TextModerationClient {
	return original.NewTextModerationClient(baseURL)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
