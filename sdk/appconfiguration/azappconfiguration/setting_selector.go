//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azappconfiguration

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/appconfiguration/azappconfiguration/internal/generated"
)

// Fields to retrieve from a configuration setting.
type SettingFields string

const (
	// The primary identifier of a configuration setting.
	SettingFieldsKey = SettingFields(generated.Enum6Key)

	// A label used to group configuration settings.
	SettingFieldsLabel = SettingFields(generated.Enum6Label)

	// The value of the configuration setting.
	SettingFieldsValue = SettingFields(generated.Enum6Value)

	// The content type of the configuration setting's value.
	SettingFieldsContentType = SettingFields(generated.Enum6ContentType)

	// An ETag indicating the version of a configuration setting within a configuration store.
	SettingFieldsETag = SettingFields(generated.Enum6Etag)

	// The last time a modifying operation was performed on the given configuration setting.
	SettingFieldsLastModified = SettingFields(generated.Enum6LastModified)

	// A value indicating whether the configuration setting is read-only.
	SettingFieldsIsReadOnly = SettingFields(generated.Enum6Locked)

	// A list of tags that can help identify what a configuration setting may be applicable for.
	SettingFieldsTags = SettingFields(generated.Enum6Tags)
)

// SettingSelector is a set of options that allows selecting a filtered set of configuration setting entities
// from the configuration store, and optionally allows indicating which fields of each setting to retrieve.
type SettingSelector struct {
	KeyFilter      *string
	LabelFilter    *string
	AcceptDateTime *time.Time
	Fields         []SettingFields
}

func AllSettingFields() []SettingFields {
	return []SettingFields{
		SettingFieldsKey,
		SettingFieldsLabel,
		SettingFieldsValue,
		SettingFieldsContentType,
		SettingFieldsETag,
		SettingFieldsLastModified,
		SettingFieldsIsReadOnly,
		SettingFieldsTags,
	}
}

func (sc SettingSelector) toGenerated() *generated.AzureAppConfigurationClientGetRevisionsOptions {
	var dt *string
	if sc.AcceptDateTime != nil {
		str := sc.AcceptDateTime.Format(timeFormat)
		dt = &str
	}

	sf := make([]generated.Enum6, len(sc.Fields))
	for i := range sc.Fields {
		sf[i] = (generated.Enum6)(sc.Fields[i])
	}

	return &generated.AzureAppConfigurationClientGetRevisionsOptions{
		After:  dt,
		Key:    sc.KeyFilter,
		Label:  sc.LabelFilter,
		Select: sf,
	}
}
