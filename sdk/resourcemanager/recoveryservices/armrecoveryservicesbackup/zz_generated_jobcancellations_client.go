//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// JobCancellationsClient contains the methods for the JobCancellations group.
// Don't use this type directly, use NewJobCancellationsClient() instead.
type JobCancellationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewJobCancellationsClient creates a new instance of JobCancellationsClient with the specified values.
func NewJobCancellationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *JobCancellationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &JobCancellationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Trigger - Cancels a job. This is an asynchronous operation. To know the status of the cancellation, call GetCancelOperationResult API.
// If the operation fails it returns the *CloudError error type.
func (client *JobCancellationsClient) Trigger(ctx context.Context, vaultName string, resourceGroupName string, jobName string, options *JobCancellationsTriggerOptions) (JobCancellationsTriggerResponse, error) {
	req, err := client.triggerCreateRequest(ctx, vaultName, resourceGroupName, jobName, options)
	if err != nil {
		return JobCancellationsTriggerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobCancellationsTriggerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return JobCancellationsTriggerResponse{}, client.triggerHandleError(resp)
	}
	return JobCancellationsTriggerResponse{RawResponse: resp}, nil
}

// triggerCreateRequest creates the Trigger request.
func (client *JobCancellationsClient) triggerCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, jobName string, options *JobCancellationsTriggerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}/cancel"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// triggerHandleError handles the Trigger error response.
func (client *JobCancellationsClient) triggerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
