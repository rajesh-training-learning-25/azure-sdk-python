//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package settings

// GetSettingResponse contains the response from method Client.GetSetting.
type GetSettingResponse struct {
	Setting
}

// GetSettingsResponse contains the response from method Client.GetSettings.
type GetSettingsResponse struct {
	ListResult
}

// UpdateSettingResponse contains the response from method Client.UpdateSetting.
type UpdateSettingResponse struct {
	Setting
}
