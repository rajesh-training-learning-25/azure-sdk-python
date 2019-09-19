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

package runtime

import original "github.com/Azure/azure-sdk-for-go/services/stable/cognitiveservices/v3.0/luis/runtime"

type BaseClient = original.BaseClient
type DynamicList = original.DynamicList
type Error = original.Error
type ErrorBody = original.ErrorBody
type ExternalEntity = original.ExternalEntity
type Intent = original.Intent
type Prediction = original.Prediction
type PredictionClient = original.PredictionClient
type PredictionRequest = original.PredictionRequest
type PredictionRequestOptions = original.PredictionRequestOptions
type PredictionResponse = original.PredictionResponse
type RequestList = original.RequestList
type Sentiment = original.Sentiment

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewPredictionClient(endpoint string) PredictionClient {
	return original.NewPredictionClient(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
