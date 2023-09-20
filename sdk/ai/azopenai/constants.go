//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

// AudioTaskLabel - Defines the possible descriptors for available audio operation responses.
type AudioTaskLabel string

const (
	AudioTaskLabelTranscribe AudioTaskLabel = "transcribe"
	AudioTaskLabelTranslate  AudioTaskLabel = "translate"
)

// PossibleAudioTaskLabelValues returns the possible values for the AudioTaskLabel const type.
func PossibleAudioTaskLabelValues() []AudioTaskLabel {
	return []AudioTaskLabel{
		AudioTaskLabelTranscribe,
		AudioTaskLabelTranslate,
	}
}

// AudioTranscriptionFormat - Defines available options for the underlying response format of output transcription information.
type AudioTranscriptionFormat string

const (
	AudioTranscriptionFormatJSON        AudioTranscriptionFormat = "json"
	AudioTranscriptionFormatSrt         AudioTranscriptionFormat = "srt"
	AudioTranscriptionFormatText        AudioTranscriptionFormat = "text"
	AudioTranscriptionFormatVerboseJSON AudioTranscriptionFormat = "verbose_json"
	AudioTranscriptionFormatVtt         AudioTranscriptionFormat = "vtt"
)

// PossibleAudioTranscriptionFormatValues returns the possible values for the AudioTranscriptionFormat const type.
func PossibleAudioTranscriptionFormatValues() []AudioTranscriptionFormat {
	return []AudioTranscriptionFormat{
		AudioTranscriptionFormatJSON,
		AudioTranscriptionFormatSrt,
		AudioTranscriptionFormatText,
		AudioTranscriptionFormatVerboseJSON,
		AudioTranscriptionFormatVtt,
	}
}

// AudioTranslationFormat - Defines available options for the underlying response format of output translation information.
type AudioTranslationFormat string

const (
	AudioTranslationFormatJSON        AudioTranslationFormat = "json"
	AudioTranslationFormatSrt         AudioTranslationFormat = "srt"
	AudioTranslationFormatText        AudioTranslationFormat = "text"
	AudioTranslationFormatVerboseJSON AudioTranslationFormat = "verbose_json"
	AudioTranslationFormatVtt         AudioTranslationFormat = "vtt"
)

// PossibleAudioTranslationFormatValues returns the possible values for the AudioTranslationFormat const type.
func PossibleAudioTranslationFormatValues() []AudioTranslationFormat {
	return []AudioTranslationFormat{
		AudioTranslationFormatJSON,
		AudioTranslationFormatSrt,
		AudioTranslationFormatText,
		AudioTranslationFormatVerboseJSON,
		AudioTranslationFormatVtt,
	}
}

// AzureChatExtensionType - A representation of configuration data for a single Azure OpenAI chat extension. This will be
// used by a chat completions request that should use Azure OpenAI chat extensions to augment the response
// behavior. The use of this configuration is compatible only with Azure OpenAI.
type AzureChatExtensionType string

const (
	// AzureChatExtensionTypeAzureCognitiveSearch enables the use of an Azure Cognitive Search index with chat completions.
	// [AzureChatExtensionConfiguration.Parameter] should be of type [AzureCognitiveSearchChatExtensionConfiguration].
	AzureChatExtensionTypeAzureCognitiveSearch AzureChatExtensionType = "AzureCognitiveSearch"
)

// PossibleAzureChatExtensionTypeValues returns the possible values for the AzureChatExtensionType const type.
func PossibleAzureChatExtensionTypeValues() []AzureChatExtensionType {
	return []AzureChatExtensionType{
		AzureChatExtensionTypeAzureCognitiveSearch,
	}
}

// AzureCognitiveSearchChatExtensionConfigurationType - The type label to use when configuring Azure OpenAI chat extensions.
// This should typically not be changed from its default value for Azure Cognitive Search.
type AzureCognitiveSearchChatExtensionConfigurationType string

const (
	AzureCognitiveSearchChatExtensionConfigurationTypeAzureCognitiveSearch AzureCognitiveSearchChatExtensionConfigurationType = "AzureCognitiveSearch"
)

// PossibleAzureCognitiveSearchChatExtensionConfigurationTypeValues returns the possible values for the AzureCognitiveSearchChatExtensionConfigurationType const type.
func PossibleAzureCognitiveSearchChatExtensionConfigurationTypeValues() []AzureCognitiveSearchChatExtensionConfigurationType {
	return []AzureCognitiveSearchChatExtensionConfigurationType{
		AzureCognitiveSearchChatExtensionConfigurationTypeAzureCognitiveSearch,
	}
}

// AzureCognitiveSearchQueryType - The type of Azure Cognitive Search retrieval query that should be executed when using it
// as an Azure OpenAI chat extension.
type AzureCognitiveSearchQueryType string

const (
	AzureCognitiveSearchQueryTypeSemantic             AzureCognitiveSearchQueryType = "semantic"
	AzureCognitiveSearchQueryTypeSimple               AzureCognitiveSearchQueryType = "simple"
	AzureCognitiveSearchQueryTypeVector               AzureCognitiveSearchQueryType = "vector"
	AzureCognitiveSearchQueryTypeVectorSemanticHybrid AzureCognitiveSearchQueryType = "vectorSemanticHybrid"
	AzureCognitiveSearchQueryTypeVectorSimpleHybrid   AzureCognitiveSearchQueryType = "vectorSimpleHybrid"
)

// PossibleAzureCognitiveSearchQueryTypeValues returns the possible values for the AzureCognitiveSearchQueryType const type.
func PossibleAzureCognitiveSearchQueryTypeValues() []AzureCognitiveSearchQueryType {
	return []AzureCognitiveSearchQueryType{
		AzureCognitiveSearchQueryTypeSemantic,
		AzureCognitiveSearchQueryTypeSimple,
		AzureCognitiveSearchQueryTypeVector,
		AzureCognitiveSearchQueryTypeVectorSemanticHybrid,
		AzureCognitiveSearchQueryTypeVectorSimpleHybrid,
	}
}

// azureOpenAIOperationState - The state of a job or item.
type azureOpenAIOperationState string

const (
	azureOpenAIOperationStateCanceled   azureOpenAIOperationState = "canceled"
	azureOpenAIOperationStateFailed     azureOpenAIOperationState = "failed"
	azureOpenAIOperationStateNotRunning azureOpenAIOperationState = "notRunning"
	azureOpenAIOperationStateRunning    azureOpenAIOperationState = "running"
	azureOpenAIOperationStateSucceeded  azureOpenAIOperationState = "succeeded"
)

// ChatRole - A description of the intended purpose of a message within a chat completions interaction.
type ChatRole string

const (
	ChatRoleAssistant ChatRole = "assistant"
	ChatRoleFunction  ChatRole = "function"
	ChatRoleSystem    ChatRole = "system"
	ChatRoleTool      ChatRole = "tool"
	ChatRoleUser      ChatRole = "user"
)

// PossibleChatRoleValues returns the possible values for the ChatRole const type.
func PossibleChatRoleValues() []ChatRole {
	return []ChatRole{
		ChatRoleAssistant,
		ChatRoleFunction,
		ChatRoleSystem,
		ChatRoleTool,
		ChatRoleUser,
	}
}

// CompletionsFinishReason - Representation of the manner in which a completions response concluded.
type CompletionsFinishReason string

const (
	CompletionsFinishReasonContentFilter CompletionsFinishReason = "content_filter"
	CompletionsFinishReasonFunctionCall  CompletionsFinishReason = "function_call"
	CompletionsFinishReasonLength        CompletionsFinishReason = "length"
	CompletionsFinishReasonStop          CompletionsFinishReason = "stop"
)

// PossibleCompletionsFinishReasonValues returns the possible values for the CompletionsFinishReason const type.
func PossibleCompletionsFinishReasonValues() []CompletionsFinishReason {
	return []CompletionsFinishReason{
		CompletionsFinishReasonContentFilter,
		CompletionsFinishReasonFunctionCall,
		CompletionsFinishReasonLength,
		CompletionsFinishReasonStop,
	}
}

// ContentFilterSeverity - Ratings for the intensity and risk level of harmful content.
type ContentFilterSeverity string

const (
	ContentFilterSeverityHigh   ContentFilterSeverity = "high"
	ContentFilterSeverityLow    ContentFilterSeverity = "low"
	ContentFilterSeverityMedium ContentFilterSeverity = "medium"
	ContentFilterSeveritySafe   ContentFilterSeverity = "safe"
)

// PossibleContentFilterSeverityValues returns the possible values for the ContentFilterSeverity const type.
func PossibleContentFilterSeverityValues() []ContentFilterSeverity {
	return []ContentFilterSeverity{
		ContentFilterSeverityHigh,
		ContentFilterSeverityLow,
		ContentFilterSeverityMedium,
		ContentFilterSeveritySafe,
	}
}

// FunctionCallPreset - The collection of predefined behaviors for handling request-provided function information in a chat
// completions operation.
type FunctionCallPreset string

const (
	FunctionCallPresetAuto FunctionCallPreset = "auto"
	FunctionCallPresetNone FunctionCallPreset = "none"
)

// PossibleFunctionCallPresetValues returns the possible values for the FunctionCallPreset const type.
func PossibleFunctionCallPresetValues() []FunctionCallPreset {
	return []FunctionCallPreset{
		FunctionCallPresetAuto,
		FunctionCallPresetNone,
	}
}

// ImageGenerationResponseFormat - The format in which the generated images are returned.
type ImageGenerationResponseFormat string

const (
	ImageGenerationResponseFormatB64JSON ImageGenerationResponseFormat = "b64_json"
	ImageGenerationResponseFormatURL     ImageGenerationResponseFormat = "url"
)

// PossibleImageGenerationResponseFormatValues returns the possible values for the ImageGenerationResponseFormat const type.
func PossibleImageGenerationResponseFormatValues() []ImageGenerationResponseFormat {
	return []ImageGenerationResponseFormat{
		ImageGenerationResponseFormatB64JSON,
		ImageGenerationResponseFormatURL,
	}
}

// ImageSize - The desired size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
type ImageSize string

const (
	ImageSize512x512   ImageSize = "512x512"
	ImageSize1024x1024 ImageSize = "1024x1024"
	ImageSize256x256   ImageSize = "256x256"
)

// PossibleImageSizeValues returns the possible values for the ImageSize const type.
func PossibleImageSizeValues() []ImageSize {
	return []ImageSize{
		ImageSize512x512,
		ImageSize1024x1024,
		ImageSize256x256,
	}
}
