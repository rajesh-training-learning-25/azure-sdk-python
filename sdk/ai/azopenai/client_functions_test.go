// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

type Params struct {
	Type       string                   `json:"type"`
	Properties map[string]ParamProperty `json:"properties"`
	Required   []string                 `json:"required,omitempty"`
}

type ParamProperty struct {
	Type        string   `json:"type"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}

func TestGetChatCompletions_usingFunctions(t *testing.T) {
	// https://platform.openai.com/docs/guides/gpt/function-calling

	t.Run("OpenAI", func(t *testing.T) {
		chatClient := newOpenAIClientForTest(t)
		testChatCompletionsFunctions(t, chatClient, openAI)
	})

	t.Run("AzureOpenAI", func(t *testing.T) {
		chatClient := newAzureOpenAIClientForTest(t, azureOpenAI)
		testChatCompletionsFunctions(t, chatClient, azureOpenAI)
	})
}

func TestGetChatCompletions_usingFunctions_streaming(t *testing.T) {
	// https://platform.openai.com/docs/guides/gpt/function-calling

	t.Run("OpenAI", func(t *testing.T) {
		chatClient := newOpenAIClientForTest(t)
		testChatCompletionsFunctionsStreaming(t, chatClient, openAI)
	})

	t.Run("AzureOpenAI", func(t *testing.T) {
		chatClient := newAzureOpenAIClientForTest(t, azureOpenAI)
		testChatCompletionsFunctionsStreaming(t, chatClient, azureOpenAI)
	})
}

func testChatCompletionsFunctions(t *testing.T, chatClient *azopenai.Client, tv testVars) {
	body := azopenai.ChatCompletionsOptions{
		DeploymentName: &tv.ChatCompletions,
		Messages: []azopenai.ChatRequestMessageClassification{
			&azopenai.ChatRequestAssistantMessage{
				Content: to.Ptr("What's the weather like in Boston, MA, in celsius?"),
			},
		},
		Tools: []azopenai.ChatCompletionsToolDefinitionClassification{
			&azopenai.ChatCompletionsFunctionToolDefinition{
				Function: &azopenai.FunctionDefinition{
					Name:        to.Ptr("get_current_weather"),
					Description: to.Ptr("Get the current weather in a given location"),
					Parameters: Params{
						Required: []string{"location"},
						Type:     "object",
						Properties: map[string]ParamProperty{
							"location": {
								Type:        "string",
								Description: "The city and state, e.g. San Francisco, CA",
							},
							"unit": {
								Type: "string",
								Enum: []string{"celsius", "fahrenheit"},
							},
						},
					},
				},
			},
		},
		Temperature: to.Ptr[float32](0.0),
	}

	resp, err := chatClient.GetChatCompletions(context.Background(), body, nil)
	require.NoError(t, err)

	funcCall := resp.Choices[0].Message.ToolCalls[0].(*azopenai.ChatCompletionsFunctionToolCall).Function

	require.Equal(t, "get_current_weather", *funcCall.Name)

	type location struct {
		Location string `json:"location"`
		Unit     string `json:"unit"`
	}

	var funcParams *location
	err = json.Unmarshal([]byte(*funcCall.Arguments), &funcParams)
	require.NoError(t, err)

	require.Equal(t, location{Location: "Boston, MA", Unit: "celsius"}, *funcParams)
}

func testChatCompletionsFunctionsStreaming(t *testing.T, chatClient *azopenai.Client, tv testVars) {
	body := azopenai.ChatCompletionsOptions{
		DeploymentName: &tv.ChatCompletions,
		Messages: []azopenai.ChatRequestMessageClassification{
			&azopenai.ChatRequestAssistantMessage{
				Content: to.Ptr("What's the weather like in Boston, MA, in celsius?"),
			},
		},
		Tools: []azopenai.ChatCompletionsToolDefinitionClassification{
			&azopenai.ChatCompletionsFunctionToolDefinition{
				Function: &azopenai.FunctionDefinition{
					Name:        to.Ptr("get_current_weather"),
					Description: to.Ptr("Get the current weather in a given location"),
					Parameters: Params{
						Required: []string{"location"},
						Type:     "object",
						Properties: map[string]ParamProperty{
							"location": {
								Type:        "string",
								Description: "The city and state, e.g. San Francisco, CA",
							},
							"unit": {
								Type: "string",
								Enum: []string{"celsius", "fahrenheit"},
							},
						},
					},
				},
			},
		},
		Temperature: to.Ptr[float32](0.0),
	}

	resp, err := chatClient.GetChatCompletionsStream(context.Background(), body, nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)

	defer func() {
		err := resp.ChatCompletionsStream.Close()
		require.NoError(t, err)
	}()

	// these results are way trickier than they should be, but we have to accumulate across
	// multiple fields to get a full result.

	funcCall := &azopenai.FunctionCall{
		Arguments: to.Ptr(""),
		Name:      to.Ptr(""),
	}

	for {
		streamResp, err := resp.ChatCompletionsStream.Read()
		require.NoError(t, err)

		if len(streamResp.Choices) == 0 {
			// there are prompt filter results.
			require.NotEmpty(t, streamResp.PromptFilterResults)
			continue
		}

		if streamResp.Choices[0].FinishReason != nil {
			break
		}

		var functionToolCall *azopenai.ChatCompletionsFunctionToolCall = streamResp.Choices[0].Delta.ToolCalls[0].(*azopenai.ChatCompletionsFunctionToolCall)
		require.NotEmpty(t, functionToolCall.Function)

		if functionToolCall.Function.Arguments != nil {
			*funcCall.Arguments += *functionToolCall.Function.Arguments
		}

		if functionToolCall.Function.Name != nil {
			*funcCall.Name += *functionToolCall.Function.Name
		}
	}

	require.Equal(t, "get_current_weather", *funcCall.Name)

	type location struct {
		Location string `json:"location"`
		Unit     string `json:"unit"`
	}

	var funcParams *location
	err = json.Unmarshal([]byte(*funcCall.Arguments), &funcParams)
	require.NoError(t, err)

	require.Equal(t, location{Location: "Boston, MA", Unit: "celsius"}, *funcParams)
}
