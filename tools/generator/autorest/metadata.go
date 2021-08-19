// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package autorest

import (
	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest/model"
	"strings"
)

func GetAdditionalOptions(metadata GenerationMetadata) model.Options {
	// additional options
	additionalOptions, _ := model.ParseOptions(strings.Split(metadata.AdditionalProperties.AdditionalOptions, " "))

	return additionalOptions
}