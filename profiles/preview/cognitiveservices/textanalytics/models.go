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

package textanalytics

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.1/textanalytics"

type BaseClient = original.BaseClient
type BatchInput = original.BatchInput
type DetectedLanguage = original.DetectedLanguage
type EntitiesBatchResultItemV2dot1 = original.EntitiesBatchResultItemV2dot1
type EntitiesBatchResultV2dot1 = original.EntitiesBatchResultV2dot1
type EntityRecordV2dot1 = original.EntityRecordV2dot1
type ErrorRecord = original.ErrorRecord
type ErrorResponse = original.ErrorResponse
type Input = original.Input
type InternalError = original.InternalError
type KeyPhraseBatchResult = original.KeyPhraseBatchResult
type KeyPhraseBatchResultItem = original.KeyPhraseBatchResultItem
type LanguageBatchResult = original.LanguageBatchResult
type LanguageBatchResultItem = original.LanguageBatchResultItem
type MatchRecordV2dot1 = original.MatchRecordV2dot1
type MultiLanguageBatchInput = original.MultiLanguageBatchInput
type MultiLanguageInput = original.MultiLanguageInput
type SentimentBatchResult = original.SentimentBatchResult
type SentimentBatchResultItem = original.SentimentBatchResultItem

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
