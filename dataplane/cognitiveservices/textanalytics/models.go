package textanalytics

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.22.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// AzureRegions enumerates the values for azure regions.
type AzureRegions string

const (
	// Eastus2 specifies the eastus 2 state for azure regions.
	Eastus2 AzureRegions = "eastus2"
	// Southeastasia specifies the southeastasia state for azure regions.
	Southeastasia AzureRegions = "southeastasia"
	// Westcentralus specifies the westcentralus state for azure regions.
	Westcentralus AzureRegions = "westcentralus"
	// Westeurope specifies the westeurope state for azure regions.
	Westeurope AzureRegions = "westeurope"
	// Westus specifies the westus state for azure regions.
	Westus AzureRegions = "westus"
)

// BatchInput is
type BatchInput struct {
	Documents *[]Input `json:"documents,omitempty"`
}

// DetectedLanguage is
type DetectedLanguage struct {
	Name        *string  `json:"name,omitempty"`
	Iso6391Name *string  `json:"iso6391Name,omitempty"`
	Score       *float64 `json:"score,omitempty"`
}

// ErrorRecord is
type ErrorRecord struct {
	ID      *string `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ErrorResponse is
type ErrorResponse struct {
	Code       *string        `json:"code,omitempty"`
	Message    *string        `json:"message,omitempty"`
	Target     *string        `json:"target,omitempty"`
	InnerError *InternalError `json:"innerError,omitempty"`
}

// Input is
type Input struct {
	ID   *string `json:"id,omitempty"`
	Text *string `json:"text,omitempty"`
}

// InternalError is
type InternalError struct {
	Code       *string        `json:"code,omitempty"`
	Message    *string        `json:"message,omitempty"`
	InnerError *InternalError `json:"innerError,omitempty"`
}

// KeyPhraseBatchResult is
type KeyPhraseBatchResult struct {
	autorest.Response `json:"-"`
	Documents         *[]KeyPhraseBatchResultItem `json:"documents,omitempty"`
	Errors            *[]ErrorRecord              `json:"errors,omitempty"`
}

// KeyPhraseBatchResultItem is
type KeyPhraseBatchResultItem struct {
	KeyPhrases *[]string `json:"keyPhrases,omitempty"`
	ID         *string   `json:"id,omitempty"`
}

// LanguageBatchResult is
type LanguageBatchResult struct {
	autorest.Response `json:"-"`
	Documents         *[]LanguageBatchResultItem `json:"documents,omitempty"`
	Errors            *[]ErrorRecord             `json:"errors,omitempty"`
}

// LanguageBatchResultItem is
type LanguageBatchResultItem struct {
	ID                *string             `json:"id,omitempty"`
	DetectedLanguages *[]DetectedLanguage `json:"detectedLanguages,omitempty"`
}

// MultiLanguageBatchInput is
type MultiLanguageBatchInput struct {
	Documents *[]MultiLanguageInput `json:"documents,omitempty"`
}

// MultiLanguageInput is
type MultiLanguageInput struct {
	Language *string `json:"language,omitempty"`
	ID       *string `json:"id,omitempty"`
	Text     *string `json:"text,omitempty"`
}

// SentimentBatchResult is
type SentimentBatchResult struct {
	autorest.Response `json:"-"`
	Documents         *[]SentimentBatchResultItem `json:"documents,omitempty"`
	Errors            *[]ErrorRecord              `json:"errors,omitempty"`
}

// SentimentBatchResultItem is
type SentimentBatchResultItem struct {
	Score *float64 `json:"score,omitempty"`
	ID    *string  `json:"id,omitempty"`
}
