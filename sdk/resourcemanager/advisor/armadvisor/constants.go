//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	moduleVersion = "v1.2.0"
)

// CPUThreshold - Minimum percentage threshold for Advisor low CPU utilization evaluation. Valid only for subscriptions. Valid
// values: 5 (default), 10, 15 or 20.
type CPUThreshold string

const (
	CPUThresholdFifteen CPUThreshold = "15"
	CPUThresholdFive    CPUThreshold = "5"
	CPUThresholdTen     CPUThreshold = "10"
	CPUThresholdTwenty  CPUThreshold = "20"
)

// PossibleCPUThresholdValues returns the possible values for the CPUThreshold const type.
func PossibleCPUThresholdValues() []CPUThreshold {
	return []CPUThreshold{
		CPUThresholdFifteen,
		CPUThresholdFive,
		CPUThresholdTen,
		CPUThresholdTwenty,
	}
}

type Category string

const (
	CategoryCost                  Category = "Cost"
	CategoryHighAvailability      Category = "HighAvailability"
	CategoryOperationalExcellence Category = "OperationalExcellence"
	CategoryPerformance           Category = "Performance"
	CategorySecurity              Category = "Security"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryCost,
		CategoryHighAvailability,
		CategoryOperationalExcellence,
		CategoryPerformance,
		CategorySecurity,
	}
}

type ConfigurationName string

const (
	ConfigurationNameDefault ConfigurationName = "default"
)

// PossibleConfigurationNameValues returns the possible values for the ConfigurationName const type.
func PossibleConfigurationNameValues() []ConfigurationName {
	return []ConfigurationName{
		ConfigurationNameDefault,
	}
}

// DigestConfigState - State of digest configuration.
type DigestConfigState string

const (
	DigestConfigStateActive   DigestConfigState = "Active"
	DigestConfigStateDisabled DigestConfigState = "Disabled"
)

// PossibleDigestConfigStateValues returns the possible values for the DigestConfigState const type.
func PossibleDigestConfigStateValues() []DigestConfigState {
	return []DigestConfigState{
		DigestConfigStateActive,
		DigestConfigStateDisabled,
	}
}

// Impact - The business impact of the recommendation.
type Impact string

const (
	ImpactHigh   Impact = "High"
	ImpactLow    Impact = "Low"
	ImpactMedium Impact = "Medium"
)

// PossibleImpactValues returns the possible values for the Impact const type.
func PossibleImpactValues() []Impact {
	return []Impact{
		ImpactHigh,
		ImpactLow,
		ImpactMedium,
	}
}

// Risk - The potential risk of not implementing the recommendation.
type Risk string

const (
	RiskError   Risk = "Error"
	RiskNone    Risk = "None"
	RiskWarning Risk = "Warning"
)

// PossibleRiskValues returns the possible values for the Risk const type.
func PossibleRiskValues() []Risk {
	return []Risk{
		RiskError,
		RiskNone,
		RiskWarning,
	}
}

type Scenario string

const (
	ScenarioAlerts Scenario = "Alerts"
)

// PossibleScenarioValues returns the possible values for the Scenario const type.
func PossibleScenarioValues() []Scenario {
	return []Scenario{
		ScenarioAlerts,
	}
}
