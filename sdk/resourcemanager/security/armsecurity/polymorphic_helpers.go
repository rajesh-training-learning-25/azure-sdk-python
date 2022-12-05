//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import "encoding/json"

func unmarshalAdditionalDataClassification(rawMsg json.RawMessage) (AdditionalDataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AdditionalDataClassification
	switch m["assessedResourceType"] {
	case string(AssessedResourceTypeContainerRegistryVulnerability):
		b = &ContainerRegistryVulnerabilityProperties{}
	case "ServerVulnerabilityAssessment":
		b = &ServerVulnerabilityProperties{}
	case string(AssessedResourceTypeSQLServerVulnerability):
		b = &SQLServerVulnerabilityProperties{}
	default:
		b = &AdditionalData{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAlertSimulatorRequestPropertiesClassification(rawMsg json.RawMessage) (AlertSimulatorRequestPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AlertSimulatorRequestPropertiesClassification
	switch m["kind"] {
	case string(KindBundles):
		b = &AlertSimulatorBundlesRequestProperties{}
	default:
		b = &AlertSimulatorRequestProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAllowlistCustomAlertRuleClassification(rawMsg json.RawMessage) (AllowlistCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AllowlistCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ConnectionFromIpNotAllowed":
		b = &ConnectionFromIPNotAllowed{}
	case "ConnectionToIpNotAllowed":
		b = &ConnectionToIPNotAllowed{}
	case "LocalUserNotAllowed":
		b = &LocalUserNotAllowed{}
	case "ProcessNotAllowed":
		b = &ProcessNotAllowed{}
	default:
		b = &AllowlistCustomAlertRule{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAllowlistCustomAlertRuleClassificationArray(rawMsg json.RawMessage) ([]AllowlistCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AllowlistCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAllowlistCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAuthenticationDetailsPropertiesClassification(rawMsg json.RawMessage) (AuthenticationDetailsPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AuthenticationDetailsPropertiesClassification
	switch m["authenticationType"] {
	case string(AuthenticationTypeAwsAssumeRole):
		b = &AwAssumeRoleAuthenticationDetailsProperties{}
	case string(AuthenticationTypeAwsCreds):
		b = &AwsCredsAuthenticationDetailsProperties{}
	case string(AuthenticationTypeGcpCredentials):
		b = &GcpCredentialsDetailsProperties{}
	default:
		b = &AuthenticationDetailsProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAutomationActionClassification(rawMsg json.RawMessage) (AutomationActionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AutomationActionClassification
	switch m["actionType"] {
	case string(ActionTypeEventHub):
		b = &AutomationActionEventHub{}
	case string(ActionTypeLogicApp):
		b = &AutomationActionLogicApp{}
	case string(ActionTypeWorkspace):
		b = &AutomationActionWorkspace{}
	default:
		b = &AutomationAction{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAutomationActionClassificationArray(rawMsg json.RawMessage) ([]AutomationActionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AutomationActionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAutomationActionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAwsOrganizationalDataClassification(rawMsg json.RawMessage) (AwsOrganizationalDataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AwsOrganizationalDataClassification
	switch m["organizationMembershipType"] {
	case string(OrganizationMembershipTypeMember):
		b = &AwsOrganizationalDataMember{}
	case string(OrganizationMembershipTypeOrganization):
		b = &AwsOrganizationalDataMaster{}
	default:
		b = &AwsOrganizationalData{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalCloudOfferingClassification(rawMsg json.RawMessage) (CloudOfferingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CloudOfferingClassification
	switch m["offeringType"] {
	case string(OfferingTypeCspmMonitorAws):
		b = &CspmMonitorAwsOffering{}
	case string(OfferingTypeCspmMonitorAzureDevOps):
		b = &CspmMonitorAzureDevOpsOffering{}
	case string(OfferingTypeCspmMonitorGcp):
		b = &CspmMonitorGcpOffering{}
	case string(OfferingTypeCspmMonitorGithub):
		b = &CspmMonitorGithubOffering{}
	case string(OfferingTypeDefenderCspmAws):
		b = &DefenderCspmAwsOffering{}
	case string(OfferingTypeDefenderCspmGcp):
		b = &DefenderCspmGcpOffering{}
	case string(OfferingTypeDefenderForContainersAws):
		b = &DefenderForContainersAwsOffering{}
	case string(OfferingTypeDefenderForContainersGcp):
		b = &DefenderForContainersGcpOffering{}
	case string(OfferingTypeDefenderForDatabasesAws):
		b = &DefenderFoDatabasesAwsOffering{}
	case string(OfferingTypeDefenderForDatabasesGcp):
		b = &DefenderForDatabasesGcpOffering{}
	case string(OfferingTypeDefenderForDevOpsAzureDevOps):
		b = &DefenderForDevOpsAzureDevOpsOffering{}
	case string(OfferingTypeDefenderForDevOpsGithub):
		b = &DefenderForDevOpsGithubOffering{}
	case string(OfferingTypeDefenderForServersAws):
		b = &DefenderForServersAwsOffering{}
	case string(OfferingTypeDefenderForServersGcp):
		b = &DefenderForServersGcpOffering{}
	case string(OfferingTypeInformationProtectionAws):
		b = &InformationProtectionAwsOffering{}
	default:
		b = &CloudOffering{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalCloudOfferingClassificationArray(rawMsg json.RawMessage) ([]CloudOfferingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]CloudOfferingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalCloudOfferingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalEnvironmentDataClassification(rawMsg json.RawMessage) (EnvironmentDataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EnvironmentDataClassification
	switch m["environmentType"] {
	case string(EnvironmentTypeAwsAccount):
		b = &AwsEnvironmentData{}
	case string(EnvironmentTypeAzureDevOpsScope):
		b = &AzureDevOpsScopeEnvironmentData{}
	case string(EnvironmentTypeGcpProject):
		b = &GcpProjectEnvironmentData{}
	case string(EnvironmentTypeGithubScope):
		b = &GithubScopeEnvironmentData{}
	default:
		b = &EnvironmentData{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalGcpOrganizationalDataClassification(rawMsg json.RawMessage) (GcpOrganizationalDataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b GcpOrganizationalDataClassification
	switch m["organizationMembershipType"] {
	case string(OrganizationMembershipTypeMember):
		b = &GcpOrganizationalDataMember{}
	case string(OrganizationMembershipTypeOrganization):
		b = &GcpOrganizationalDataOrganization{}
	default:
		b = &GcpOrganizationalData{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalResourceDetailsClassification(rawMsg json.RawMessage) (ResourceDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResourceDetailsClassification
	switch m["source"] {
	case string(SourceAzure):
		b = &AzureResourceDetails{}
	case string(SourceOnPremise):
		b = &OnPremiseResourceDetails{}
	case string(SourceOnPremiseSQL):
		b = &OnPremiseSQLResourceDetails{}
	default:
		b = &ResourceDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalResourceIdentifierClassification(rawMsg json.RawMessage) (ResourceIdentifierClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResourceIdentifierClassification
	switch m["type"] {
	case string(ResourceIdentifierTypeAzureResource):
		b = &AzureResourceIdentifier{}
	case string(ResourceIdentifierTypeLogAnalytics):
		b = &LogAnalyticsIdentifier{}
	default:
		b = &ResourceIdentifier{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalResourceIdentifierClassificationArray(rawMsg json.RawMessage) ([]ResourceIdentifierClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ResourceIdentifierClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalResourceIdentifierClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalSettingClassification(rawMsg json.RawMessage) (SettingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SettingClassification
	switch m["kind"] {
	case string(SettingKindAlertSyncSettings):
		b = &AlertSyncSettings{}
	case string(SettingKindDataExportSettings):
		b = &DataExportSettings{}
	default:
		b = &Setting{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSettingClassificationArray(rawMsg json.RawMessage) ([]SettingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]SettingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalSettingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalThresholdCustomAlertRuleClassification(rawMsg json.RawMessage) (ThresholdCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ThresholdCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ActiveConnectionsNotInAllowedRange":
		b = &ActiveConnectionsNotInAllowedRange{}
	case "AmqpC2DMessagesNotInAllowedRange":
		b = &AmqpC2DMessagesNotInAllowedRange{}
	case "AmqpC2DRejectedMessagesNotInAllowedRange":
		b = &AmqpC2DRejectedMessagesNotInAllowedRange{}
	case "AmqpD2CMessagesNotInAllowedRange":
		b = &AmqpD2CMessagesNotInAllowedRange{}
	case "DirectMethodInvokesNotInAllowedRange":
		b = &DirectMethodInvokesNotInAllowedRange{}
	case "FailedLocalLoginsNotInAllowedRange":
		b = &FailedLocalLoginsNotInAllowedRange{}
	case "FileUploadsNotInAllowedRange":
		b = &FileUploadsNotInAllowedRange{}
	case "HttpC2DMessagesNotInAllowedRange":
		b = &HTTPC2DMessagesNotInAllowedRange{}
	case "HttpC2DRejectedMessagesNotInAllowedRange":
		b = &HTTPC2DRejectedMessagesNotInAllowedRange{}
	case "HttpD2CMessagesNotInAllowedRange":
		b = &HTTPD2CMessagesNotInAllowedRange{}
	case "MqttC2DMessagesNotInAllowedRange":
		b = &MqttC2DMessagesNotInAllowedRange{}
	case "MqttC2DRejectedMessagesNotInAllowedRange":
		b = &MqttC2DRejectedMessagesNotInAllowedRange{}
	case "MqttD2CMessagesNotInAllowedRange":
		b = &MqttD2CMessagesNotInAllowedRange{}
	case "QueuePurgesNotInAllowedRange":
		b = &QueuePurgesNotInAllowedRange{}
	case "TimeWindowCustomAlertRule":
		b = &TimeWindowCustomAlertRule{}
	case "TwinUpdatesNotInAllowedRange":
		b = &TwinUpdatesNotInAllowedRange{}
	case "UnauthorizedOperationsNotInAllowedRange":
		b = &UnauthorizedOperationsNotInAllowedRange{}
	default:
		b = &ThresholdCustomAlertRule{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalThresholdCustomAlertRuleClassificationArray(rawMsg json.RawMessage) ([]ThresholdCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ThresholdCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalThresholdCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTimeWindowCustomAlertRuleClassification(rawMsg json.RawMessage) (TimeWindowCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TimeWindowCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ActiveConnectionsNotInAllowedRange":
		b = &ActiveConnectionsNotInAllowedRange{}
	case "AmqpC2DMessagesNotInAllowedRange":
		b = &AmqpC2DMessagesNotInAllowedRange{}
	case "AmqpC2DRejectedMessagesNotInAllowedRange":
		b = &AmqpC2DRejectedMessagesNotInAllowedRange{}
	case "AmqpD2CMessagesNotInAllowedRange":
		b = &AmqpD2CMessagesNotInAllowedRange{}
	case "DirectMethodInvokesNotInAllowedRange":
		b = &DirectMethodInvokesNotInAllowedRange{}
	case "FailedLocalLoginsNotInAllowedRange":
		b = &FailedLocalLoginsNotInAllowedRange{}
	case "FileUploadsNotInAllowedRange":
		b = &FileUploadsNotInAllowedRange{}
	case "HttpC2DMessagesNotInAllowedRange":
		b = &HTTPC2DMessagesNotInAllowedRange{}
	case "HttpC2DRejectedMessagesNotInAllowedRange":
		b = &HTTPC2DRejectedMessagesNotInAllowedRange{}
	case "HttpD2CMessagesNotInAllowedRange":
		b = &HTTPD2CMessagesNotInAllowedRange{}
	case "MqttC2DMessagesNotInAllowedRange":
		b = &MqttC2DMessagesNotInAllowedRange{}
	case "MqttC2DRejectedMessagesNotInAllowedRange":
		b = &MqttC2DRejectedMessagesNotInAllowedRange{}
	case "MqttD2CMessagesNotInAllowedRange":
		b = &MqttD2CMessagesNotInAllowedRange{}
	case "QueuePurgesNotInAllowedRange":
		b = &QueuePurgesNotInAllowedRange{}
	case "TwinUpdatesNotInAllowedRange":
		b = &TwinUpdatesNotInAllowedRange{}
	case "UnauthorizedOperationsNotInAllowedRange":
		b = &UnauthorizedOperationsNotInAllowedRange{}
	default:
		b = &TimeWindowCustomAlertRule{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTimeWindowCustomAlertRuleClassificationArray(rawMsg json.RawMessage) ([]TimeWindowCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]TimeWindowCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTimeWindowCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}