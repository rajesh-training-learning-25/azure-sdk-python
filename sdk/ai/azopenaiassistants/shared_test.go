//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenaiassistants_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	assistants "github.com/Azure/azure-sdk-for-go/sdk/ai/azopenaiassistants"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

func stringize(v assistants.MessageContentClassification) string {
	switch m := v.(type) {
	case *assistants.MessageTextContent:
		return fmt.Sprintf("Text = %s\n", *m.Text.Value)
	case *assistants.MessageImageFileContent:
		return fmt.Sprintf("Image = %s\n", *m.ImageFile.FileID.FileID)
	}

	panic("Unhandled type for stringizing")
}

type mustGetClientWithAssistantArgs struct {
	newClientArgs
	Instructions string
}

type newClientArgs struct {
	Azure       bool
	UseIdentity bool
}

func newClient(t *testing.T, args newClientArgs) *assistants.Client {
	var httpClient policy.Transporter

	if recording.GetRecordMode() != recording.LiveMode {
		err := recording.Start(t, RecordingDirectory, nil)
		require.NoError(t, err)

		t.Cleanup(func() {
			err := recording.Stop(t, nil)
			require.NoError(t, err)
		})

		tmpHttpClient, err := recording.NewRecordingHTTPClient(t, nil)
		require.NoError(t, err)

		err = recording.AddURISanitizer("endoint", strings.TrimRight(tv.AOAIEndpoint, "/"), nil)
		require.NoError(t, err)

		err = recording.AddURISanitizer("endoint", strings.TrimRight(tv.OpenAIEndpoint, "/"), nil)
		require.NoError(t, err)

		err = recording.AddHeaderRegexSanitizer("Api-Key", "apikey", "", nil)
		require.NoError(t, err)

		httpClient = tmpHttpClient
	} else if os.Getenv("SSLKEYLOGFILE") != "" {
		file, err := os.OpenFile(os.Getenv("SSLKEYLOGFILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)
		require.NoError(t, err)

		t.Cleanup(func() {
			err := file.Close()
			require.NoError(t, err)
		})

		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.TLSClientConfig = &tls.Config{
			KeyLogWriter: file,
		}
		httpClient = &http.Client{Transport: transport}
	}

	opts := &assistants.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Logging: policy.LogOptions{
				IncludeBody: true,
			},
			Transport: httpClient,
		},
	}

	if args.Azure {
		if args.UseIdentity {
			dac, err := azidentity.NewDefaultAzureCredential(nil)
			require.NoError(t, err)

			tmpClient, err := assistants.NewClient(tv.AOAIEndpoint, dac, opts)
			require.NoError(t, err)
			return tmpClient
		} else {
			tmpClient, err := assistants.NewClientWithKeyCredential(tv.AOAIEndpoint, azcore.NewKeyCredential(tv.AOAIKey), opts)
			require.NoError(t, err)
			return tmpClient
		}
	} else {
		tmpClient, err := assistants.NewClientForOpenAI(tv.OpenAIEndpoint, azcore.NewKeyCredential(tv.OpenAIKey), opts)
		require.NoError(t, err)
		return tmpClient
	}
}

func mustGetClientWithAssistant(t *testing.T, args mustGetClientWithAssistantArgs) (*assistants.Client, assistants.CreateAssistantResponse) {
	client := newClient(t, args.newClientArgs)

	// give the assistant a random-ish name.
	assistantName := fmt.Sprintf("go-assistant-%d", time.Now().UnixNano())

	createResp, err := client.CreateAssistant(context.Background(), assistants.AssistantCreationBody{
		Name:           &assistantName,
		DeploymentName: &assistantsModel,
		Instructions:   to.Ptr("You are a personal math tutor. Write and run code to answer math questions."),
		Tools: []assistants.ToolDefinitionClassification{
			&assistants.CodeInterpreterToolDefinition{},

			// others...
			// &assistants.FunctionToolDefinition{}
			// &assistants.RetrievalToolDefinition{}
		},
	}, nil)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err := client.DeleteAssistant(context.Background(), *createResp.ID, nil)
		require.NoError(t, err)
	})

	return client, createResp
}

func mustUploadFile(t *testing.T, c *assistants.Client, text string) assistants.UploadFileResponse {
	textBytes := []byte(text)

	uploadResp, err := c.UploadFile(context.Background(), textBytes, assistants.FilePurposeAssistants, &assistants.UploadFileOptions{
		Filename: to.Ptr("a.txt"),
	})
	require.NoError(t, err)
	require.Equal(t, len(textBytes), int(*uploadResp.Bytes))

	t.Cleanup(func() {
		_, err := c.DeleteFile(context.Background(), *uploadResp.ID, nil)
		require.NoError(t, err)
	})

	return uploadResp
}
