package anomalydetector

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/anomalydetector"

// ChangePointDetectRequest ...
type ChangePointDetectRequest struct {
	// Series - Time series data points. Points should be sorted by timestamp in ascending order to match the change point detection result.
	Series *[]TimeSeriesPoint `json:"series,omitempty"`
	// Granularity - Can only be one of yearly, monthly, weekly, daily, hourly, minutely or secondly. Granularity is used for verify whether input series is valid. Possible values include: 'Yearly', 'Monthly', 'Weekly', 'Daily', 'Hourly', 'PerMinute', 'PerSecond', 'Microsecond', 'None'
	Granularity TimeGranularity `json:"granularity,omitempty"`
	// CustomInterval - Custom Interval is used to set non-standard time interval, for example, if the series is 5 minutes, request can be set as {"granularity":"minutely", "customInterval":5}.
	CustomInterval *int32 `json:"customInterval,omitempty"`
	// Period - Optional argument, periodic value of a time series. If the value is null or does not present, the API will determine the period automatically.
	Period *int32 `json:"period,omitempty"`
	// StableTrendWindow - Optional argument, advanced model parameter, a default stableTrendWindow will be used in detection.
	StableTrendWindow *int32 `json:"stableTrendWindow,omitempty"`
	// Threshold - Optional argument, advanced model parameter, between 0.0-1.0, the lower the value is, the larger the trend error will be which means less change point will be accepted.
	Threshold *float64 `json:"threshold,omitempty"`
}

// ChangePointDetectResponse ...
type ChangePointDetectResponse struct {
	autorest.Response `json:"-"`
	// Period - READ-ONLY; Frequency extracted from the series, zero means no recurrent pattern has been found.
	Period *int32 `json:"period,omitempty"`
	// IsChangePoint - isChangePoint contains change point properties for each input point. True means an anomaly either negative or positive has been detected. The index of the array is consistent with the input series.
	IsChangePoint *[]bool `json:"isChangePoint,omitempty"`
	// ConfidenceScores - the change point confidence of each point
	ConfidenceScores *[]float64 `json:"confidenceScores,omitempty"`
}

// MarshalJSON is the custom marshaler for ChangePointDetectResponse.
func (cpdr ChangePointDetectResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cpdr.IsChangePoint != nil {
		objectMap["isChangePoint"] = cpdr.IsChangePoint
	}
	if cpdr.ConfidenceScores != nil {
		objectMap["confidenceScores"] = cpdr.ConfidenceScores
	}
	return json.Marshal(objectMap)
}

// DetectRequest ...
type DetectRequest struct {
	// Series - Time series data points. Points should be sorted by timestamp in ascending order to match the anomaly detection result. If the data is not sorted correctly or there is duplicated timestamp, the API will not work. In such case, an error message will be returned.
	Series *[]TimeSeriesPoint `json:"series,omitempty"`
	// Granularity - Possible values include: 'Yearly', 'Monthly', 'Weekly', 'Daily', 'Hourly', 'PerMinute', 'PerSecond', 'Microsecond', 'None'
	Granularity TimeGranularity `json:"granularity,omitempty"`
	// CustomInterval - Custom Interval is used to set non-standard time interval, for example, if the series is 5 minutes, request can be set as {"granularity":"minutely", "customInterval":5}.
	CustomInterval *int32 `json:"customInterval,omitempty"`
	// Period - Optional argument, periodic value of a time series. If the value is null or does not present, the API will determine the period automatically.
	Period *int32 `json:"period,omitempty"`
	// MaxAnomalyRatio - Optional argument, advanced model parameter, max anomaly ratio in a time series.
	MaxAnomalyRatio *float64 `json:"maxAnomalyRatio,omitempty"`
	// Sensitivity - Optional argument, advanced model parameter, between 0-99, the lower the value is, the larger the margin value will be which means less anomalies will be accepted.
	Sensitivity *int32 `json:"sensitivity,omitempty"`
}

// EntireDetectResponse ...
type EntireDetectResponse struct {
	autorest.Response `json:"-"`
	// Period - Frequency extracted from the series, zero means no recurrent pattern has been found.
	Period *int32 `json:"period,omitempty"`
	// ExpectedValues - ExpectedValues contain expected value for each input point. The index of the array is consistent with the input series.
	ExpectedValues *[]float64 `json:"expectedValues,omitempty"`
	// UpperMargins - UpperMargins contain upper margin of each input point. UpperMargin is used to calculate upperBoundary, which equals to expectedValue + (100 - marginScale)*upperMargin. Anomalies in response can be filtered by upperBoundary and lowerBoundary. By adjusting marginScale value, less significant anomalies can be filtered in client side. The index of the array is consistent with the input series.
	UpperMargins *[]float64 `json:"upperMargins,omitempty"`
	// LowerMargins - LowerMargins contain lower margin of each input point. LowerMargin is used to calculate lowerBoundary, which equals to expectedValue - (100 - marginScale)*lowerMargin. Points between the boundary can be marked as normal ones in client side. The index of the array is consistent with the input series.
	LowerMargins *[]float64 `json:"lowerMargins,omitempty"`
	// IsAnomaly - IsAnomaly contains anomaly properties for each input point. True means an anomaly either negative or positive has been detected. The index of the array is consistent with the input series.
	IsAnomaly *[]bool `json:"isAnomaly,omitempty"`
	// IsNegativeAnomaly - IsNegativeAnomaly contains anomaly status in negative direction for each input point. True means a negative anomaly has been detected. A negative anomaly means the point is detected as an anomaly and its real value is smaller than the expected one. The index of the array is consistent with the input series.
	IsNegativeAnomaly *[]bool `json:"isNegativeAnomaly,omitempty"`
	// IsPositiveAnomaly - IsPositiveAnomaly contain anomaly status in positive direction for each input point. True means a positive anomaly has been detected. A positive anomaly means the point is detected as an anomaly and its real value is larger than the expected one. The index of the array is consistent with the input series.
	IsPositiveAnomaly *[]bool `json:"isPositiveAnomaly,omitempty"`
}

// Error error information returned by the API.
type Error struct {
	// Code - The error code.
	Code interface{} `json:"code,omitempty"`
	// Message - A message explaining the error reported by the service.
	Message *string `json:"message,omitempty"`
}

// LastDetectResponse ...
type LastDetectResponse struct {
	autorest.Response `json:"-"`
	// Period - Frequency extracted from the series, zero means no recurrent pattern has been found.
	Period *int32 `json:"period,omitempty"`
	// SuggestedWindow - Suggested input series points needed for detecting the latest point.
	SuggestedWindow *int32 `json:"suggestedWindow,omitempty"`
	// ExpectedValue - Expected value of the latest point.
	ExpectedValue *float64 `json:"expectedValue,omitempty"`
	// UpperMargin - Upper margin of the latest point. UpperMargin is used to calculate upperBoundary, which equals to expectedValue + (100 - marginScale)*upperMargin. If the value of latest point is between upperBoundary and lowerBoundary, it should be treated as normal value. By adjusting marginScale value, anomaly status of latest point can be changed.
	UpperMargin *float64 `json:"upperMargin,omitempty"`
	// LowerMargin - Lower margin of the latest point. LowerMargin is used to calculate lowerBoundary, which equals to expectedValue - (100 - marginScale)*lowerMargin.
	LowerMargin *float64 `json:"lowerMargin,omitempty"`
	// IsAnomaly - Anomaly status of the latest point, true means the latest point is an anomaly either in negative direction or positive direction.
	IsAnomaly *bool `json:"isAnomaly,omitempty"`
	// IsNegativeAnomaly - Anomaly status in negative direction of the latest point. True means the latest point is an anomaly and its real value is smaller than the expected one.
	IsNegativeAnomaly *bool `json:"isNegativeAnomaly,omitempty"`
	// IsPositiveAnomaly - Anomaly status in positive direction of the latest point. True means the latest point is an anomaly and its real value is larger than the expected one.
	IsPositiveAnomaly *bool `json:"isPositiveAnomaly,omitempty"`
}

// TimeSeriesPoint ...
type TimeSeriesPoint struct {
	// Timestamp - Optional argument, timestamp of a data point (ISO8601 format).
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// Value - The measurement of that point, should be float.
	Value *float64 `json:"value,omitempty"`
}
