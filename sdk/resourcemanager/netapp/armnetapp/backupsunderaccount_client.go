//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BackupsUnderAccountClient contains the methods for the BackupsUnderAccount group.
// Don't use this type directly, use NewBackupsUnderAccountClient() instead.
type BackupsUnderAccountClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupsUnderAccountClient creates a new instance of BackupsUnderAccountClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupsUnderAccountClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupsUnderAccountClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupsUnderAccountClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginMigrateBackups - Migrate the backups under a NetApp account to backup vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - body - Migrate backups under an account payload supplied in the body of the operation.
//   - options - BackupsUnderAccountClientBeginMigrateBackupsOptions contains the optional parameters for the BackupsUnderAccountClient.BeginMigrateBackups
//     method.
func (client *BackupsUnderAccountClient) BeginMigrateBackups(ctx context.Context, resourceGroupName string, accountName string, body BackupsMigrationRequest, options *BackupsUnderAccountClientBeginMigrateBackupsOptions) (*runtime.Poller[BackupsUnderAccountClientMigrateBackupsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.migrateBackups(ctx, resourceGroupName, accountName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsUnderAccountClientMigrateBackupsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsUnderAccountClientMigrateBackupsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// MigrateBackups - Migrate the backups under a NetApp account to backup vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *BackupsUnderAccountClient) migrateBackups(ctx context.Context, resourceGroupName string, accountName string, body BackupsMigrationRequest, options *BackupsUnderAccountClientBeginMigrateBackupsOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsUnderAccountClient.BeginMigrateBackups"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.migrateBackupsCreateRequest(ctx, resourceGroupName, accountName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// migrateBackupsCreateRequest creates the MigrateBackups request.
func (client *BackupsUnderAccountClient) migrateBackupsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, body BackupsMigrationRequest, options *BackupsUnderAccountClientBeginMigrateBackupsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/migrateBackups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
