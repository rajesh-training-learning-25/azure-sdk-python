//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// IotHubClient contains the methods for the IotHub group.
// Don't use this type directly, use NewIotHubClient() instead.
type IotHubClient struct {
	con            *connection
	subscriptionID string
}

// NewIotHubClient creates a new instance of IotHubClient with the specified values.
func NewIotHubClient(con *connection, subscriptionID string) *IotHubClient {
	return &IotHubClient{con: con, subscriptionID: subscriptionID}
}

// BeginManualFailover - Manually initiate a failover for the IoT Hub to its secondary region. To learn more, see https://aka.ms/manualfailover
// If the operation fails it returns the *ErrorDetails error type.
func (client *IotHubClient) BeginManualFailover(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *IotHubBeginManualFailoverOptions) (IotHubManualFailoverPollerResponse, error) {
	resp, err := client.manualFailover(ctx, iotHubName, resourceGroupName, failoverInput, options)
	if err != nil {
		return IotHubManualFailoverPollerResponse{}, err
	}
	result := IotHubManualFailoverPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IotHubClient.ManualFailover", "", resp, client.con.Pipeline(), client.manualFailoverHandleError)
	if err != nil {
		return IotHubManualFailoverPollerResponse{}, err
	}
	result.Poller = &IotHubManualFailoverPoller{
		pt: pt,
	}
	return result, nil
}

// ManualFailover - Manually initiate a failover for the IoT Hub to its secondary region. To learn more, see https://aka.ms/manualfailover
// If the operation fails it returns the *ErrorDetails error type.
func (client *IotHubClient) manualFailover(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *IotHubBeginManualFailoverOptions) (*http.Response, error) {
	req, err := client.manualFailoverCreateRequest(ctx, iotHubName, resourceGroupName, failoverInput, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.manualFailoverHandleError(resp)
	}
	return resp, nil
}

// manualFailoverCreateRequest creates the ManualFailover request.
func (client *IotHubClient) manualFailoverCreateRequest(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *IotHubBeginManualFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{iotHubName}/failover"
	if iotHubName == "" {
		return nil, errors.New("parameter iotHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotHubName}", url.PathEscape(iotHubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, failoverInput)
}

// manualFailoverHandleError handles the ManualFailover error response.
func (client *IotHubClient) manualFailoverHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
