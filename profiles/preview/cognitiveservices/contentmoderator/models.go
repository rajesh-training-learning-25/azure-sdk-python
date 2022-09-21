//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package contentmoderator

import original "github.com/Azure/dev/azure-sdk-for-go/services/cognitiveservices/v1.0/contentmoderator"

type StatusEnum = original.StatusEnum

const (
	Complete    StatusEnum = original.Complete
	Pending     StatusEnum = original.Pending
	Unpublished StatusEnum = original.Unpublished
)

type Type = original.Type

const (
	TypeImage Type = original.TypeImage
	TypeText  Type = original.TypeText
)

type APIError = original.APIError
type Address = original.Address
type BaseClient = original.BaseClient
type Body = original.Body
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
type IPA = original.IPA
type Image = original.Image
type ImageAdditionalInfoItem = original.ImageAdditionalInfoItem
type ImageIds = original.ImageIds
type ImageList = original.ImageList
type ImageModerationClient = original.ImageModerationClient
type ImageURL = original.ImageURL
type Job = original.Job
type JobExecutionReportDetails = original.JobExecutionReportDetails
type JobID = original.JobID
type JobListResult = original.JobListResult
type KeyValuePair = original.KeyValuePair
type ListImageList = original.ListImageList
type ListManagementImageClient = original.ListManagementImageClient
type ListManagementImageListsClient = original.ListManagementImageListsClient
type ListManagementTermClient = original.ListManagementTermClient
type ListManagementTermListsClient = original.ListManagementTermListsClient
type ListString = original.ListString
type ListTermList = original.ListTermList
type Match = original.Match
type MatchResponse = original.MatchResponse
type OCR = original.OCR
type PII = original.PII
type Phone = original.Phone
type RefreshIndex = original.RefreshIndex
type Review = original.Review
type ReviewsClient = original.ReviewsClient
type SSN = original.SSN
type Screen = original.Screen
type Status = original.Status
type String = original.String
type Tag = original.Tag
type TermList = original.TermList
type Terms = original.Terms
type TermsData = original.TermsData
type TermsInList = original.TermsInList
type TermsPaging = original.TermsPaging
type TextModerationClient = original.TextModerationClient
type TranscriptModerationBodyItem = original.TranscriptModerationBodyItem
type TranscriptModerationBodyItemTermsItem = original.TranscriptModerationBodyItemTermsItem
type VideoFrameBodyItem = original.VideoFrameBodyItem
type VideoFrameBodyItemMetadataItem = original.VideoFrameBodyItemMetadataItem
type VideoFrameBodyItemReviewerResultTagsItem = original.VideoFrameBodyItemReviewerResultTagsItem

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewImageModerationClient(endpoint string) ImageModerationClient {
	return original.NewImageModerationClient(endpoint)
}
func NewListManagementImageClient(endpoint string) ListManagementImageClient {
	return original.NewListManagementImageClient(endpoint)
}
func NewListManagementImageListsClient(endpoint string) ListManagementImageListsClient {
	return original.NewListManagementImageListsClient(endpoint)
}
func NewListManagementTermClient(endpoint string) ListManagementTermClient {
	return original.NewListManagementTermClient(endpoint)
}
func NewListManagementTermListsClient(endpoint string) ListManagementTermListsClient {
	return original.NewListManagementTermListsClient(endpoint)
}
func NewReviewsClient(endpoint string) ReviewsClient {
	return original.NewReviewsClient(endpoint)
}
func NewTextModerationClient(endpoint string) TextModerationClient {
	return original.NewTextModerationClient(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleStatusEnumValues() []StatusEnum {
	return original.PossibleStatusEnumValues()
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
