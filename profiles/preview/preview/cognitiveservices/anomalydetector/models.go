// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package anomalydetector

import original "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0/anomalydetector"

type TimeGranularity = original.TimeGranularity

const (
	Daily     TimeGranularity = original.Daily
	Hourly    TimeGranularity = original.Hourly
	Monthly   TimeGranularity = original.Monthly
	PerMinute TimeGranularity = original.PerMinute
	PerSecond TimeGranularity = original.PerSecond
	Weekly    TimeGranularity = original.Weekly
	Yearly    TimeGranularity = original.Yearly
)

type BaseClient = original.BaseClient
type ChangePointDetectRequest = original.ChangePointDetectRequest
type ChangePointDetectResponse = original.ChangePointDetectResponse
type DetectRequest = original.DetectRequest
type EntireDetectResponse = original.EntireDetectResponse
type Error = original.Error
type LastDetectResponse = original.LastDetectResponse
type TimeSeriesPoint = original.TimeSeriesPoint

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleTimeGranularityValues() []TimeGranularity {
	return original.PossibleTimeGranularityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
