//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// PrivateZonesCreateOrUpdatePollerResponse contains the response from method PrivateZones.CreateOrUpdate.
type PrivateZonesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateZonesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateZonesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateZonesCreateOrUpdateResponse, error) {
	respType := PrivateZonesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateZone)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateZonesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PrivateZonesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateZonesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateZonesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateZonesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateZonesCreateOrUpdateResponse contains the response from method PrivateZones.CreateOrUpdate.
type PrivateZonesCreateOrUpdateResponse struct {
	PrivateZonesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateZonesCreateOrUpdateResult contains the result from method PrivateZones.CreateOrUpdate.
type PrivateZonesCreateOrUpdateResult struct {
	PrivateZone
}

// PrivateZonesDeletePollerResponse contains the response from method PrivateZones.Delete.
type PrivateZonesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateZonesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateZonesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateZonesDeleteResponse, error) {
	respType := PrivateZonesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateZonesDeletePollerResponse from the provided client and resume token.
func (l *PrivateZonesDeletePollerResponse) Resume(ctx context.Context, client *PrivateZonesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateZonesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateZonesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateZonesDeleteResponse contains the response from method PrivateZones.Delete.
type PrivateZonesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateZonesGetResponse contains the response from method PrivateZones.Get.
type PrivateZonesGetResponse struct {
	PrivateZonesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateZonesGetResult contains the result from method PrivateZones.Get.
type PrivateZonesGetResult struct {
	PrivateZone
}

// PrivateZonesListByResourceGroupResponse contains the response from method PrivateZones.ListByResourceGroup.
type PrivateZonesListByResourceGroupResponse struct {
	PrivateZonesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateZonesListByResourceGroupResult contains the result from method PrivateZones.ListByResourceGroup.
type PrivateZonesListByResourceGroupResult struct {
	PrivateZoneListResult
}

// PrivateZonesListResponse contains the response from method PrivateZones.List.
type PrivateZonesListResponse struct {
	PrivateZonesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateZonesListResult contains the result from method PrivateZones.List.
type PrivateZonesListResult struct {
	PrivateZoneListResult
}

// PrivateZonesUpdatePollerResponse contains the response from method PrivateZones.Update.
type PrivateZonesUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateZonesUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateZonesUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateZonesUpdateResponse, error) {
	respType := PrivateZonesUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateZone)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateZonesUpdatePollerResponse from the provided client and resume token.
func (l *PrivateZonesUpdatePollerResponse) Resume(ctx context.Context, client *PrivateZonesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateZonesClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateZonesUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateZonesUpdateResponse contains the response from method PrivateZones.Update.
type PrivateZonesUpdateResponse struct {
	PrivateZonesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateZonesUpdateResult contains the result from method PrivateZones.Update.
type PrivateZonesUpdateResult struct {
	PrivateZone
}

// RecordSetsCreateOrUpdateResponse contains the response from method RecordSets.CreateOrUpdate.
type RecordSetsCreateOrUpdateResponse struct {
	RecordSetsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsCreateOrUpdateResult contains the result from method RecordSets.CreateOrUpdate.
type RecordSetsCreateOrUpdateResult struct {
	RecordSet
}

// RecordSetsDeleteResponse contains the response from method RecordSets.Delete.
type RecordSetsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsGetResponse contains the response from method RecordSets.Get.
type RecordSetsGetResponse struct {
	RecordSetsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsGetResult contains the result from method RecordSets.Get.
type RecordSetsGetResult struct {
	RecordSet
}

// RecordSetsListByTypeResponse contains the response from method RecordSets.ListByType.
type RecordSetsListByTypeResponse struct {
	RecordSetsListByTypeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsListByTypeResult contains the result from method RecordSets.ListByType.
type RecordSetsListByTypeResult struct {
	RecordSetListResult
}

// RecordSetsListResponse contains the response from method RecordSets.List.
type RecordSetsListResponse struct {
	RecordSetsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsListResult contains the result from method RecordSets.List.
type RecordSetsListResult struct {
	RecordSetListResult
}

// RecordSetsUpdateResponse contains the response from method RecordSets.Update.
type RecordSetsUpdateResponse struct {
	RecordSetsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecordSetsUpdateResult contains the result from method RecordSets.Update.
type RecordSetsUpdateResult struct {
	RecordSet
}

// VirtualNetworkLinksCreateOrUpdatePollerResponse contains the response from method VirtualNetworkLinks.CreateOrUpdate.
type VirtualNetworkLinksCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VirtualNetworkLinksCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l VirtualNetworkLinksCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VirtualNetworkLinksCreateOrUpdateResponse, error) {
	respType := VirtualNetworkLinksCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.VirtualNetworkLink)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VirtualNetworkLinksCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *VirtualNetworkLinksCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *VirtualNetworkLinksClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VirtualNetworkLinksClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &VirtualNetworkLinksCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VirtualNetworkLinksCreateOrUpdateResponse contains the response from method VirtualNetworkLinks.CreateOrUpdate.
type VirtualNetworkLinksCreateOrUpdateResponse struct {
	VirtualNetworkLinksCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkLinksCreateOrUpdateResult contains the result from method VirtualNetworkLinks.CreateOrUpdate.
type VirtualNetworkLinksCreateOrUpdateResult struct {
	VirtualNetworkLink
}

// VirtualNetworkLinksDeletePollerResponse contains the response from method VirtualNetworkLinks.Delete.
type VirtualNetworkLinksDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VirtualNetworkLinksDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l VirtualNetworkLinksDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VirtualNetworkLinksDeleteResponse, error) {
	respType := VirtualNetworkLinksDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VirtualNetworkLinksDeletePollerResponse from the provided client and resume token.
func (l *VirtualNetworkLinksDeletePollerResponse) Resume(ctx context.Context, client *VirtualNetworkLinksClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VirtualNetworkLinksClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &VirtualNetworkLinksDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VirtualNetworkLinksDeleteResponse contains the response from method VirtualNetworkLinks.Delete.
type VirtualNetworkLinksDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkLinksGetResponse contains the response from method VirtualNetworkLinks.Get.
type VirtualNetworkLinksGetResponse struct {
	VirtualNetworkLinksGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkLinksGetResult contains the result from method VirtualNetworkLinks.Get.
type VirtualNetworkLinksGetResult struct {
	VirtualNetworkLink
}

// VirtualNetworkLinksListResponse contains the response from method VirtualNetworkLinks.List.
type VirtualNetworkLinksListResponse struct {
	VirtualNetworkLinksListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkLinksListResult contains the result from method VirtualNetworkLinks.List.
type VirtualNetworkLinksListResult struct {
	VirtualNetworkLinkListResult
}

// VirtualNetworkLinksUpdatePollerResponse contains the response from method VirtualNetworkLinks.Update.
type VirtualNetworkLinksUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VirtualNetworkLinksUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l VirtualNetworkLinksUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VirtualNetworkLinksUpdateResponse, error) {
	respType := VirtualNetworkLinksUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.VirtualNetworkLink)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VirtualNetworkLinksUpdatePollerResponse from the provided client and resume token.
func (l *VirtualNetworkLinksUpdatePollerResponse) Resume(ctx context.Context, client *VirtualNetworkLinksClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VirtualNetworkLinksClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &VirtualNetworkLinksUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VirtualNetworkLinksUpdateResponse contains the response from method VirtualNetworkLinks.Update.
type VirtualNetworkLinksUpdateResponse struct {
	VirtualNetworkLinksUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworkLinksUpdateResult contains the result from method VirtualNetworkLinks.Update.
type VirtualNetworkLinksUpdateResult struct {
	VirtualNetworkLink
}
