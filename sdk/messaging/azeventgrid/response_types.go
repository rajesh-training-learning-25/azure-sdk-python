//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azeventgrid

// ClientAcknowledgeCloudEventsResponse contains the response from method Client.AcknowledgeCloudEvents.
type ClientAcknowledgeCloudEventsResponse struct {
	AcknowledgeResult
}

// ClientPublishCloudEventsResponse contains the response from method Client.PublishCloudEvents.
type ClientPublishCloudEventsResponse struct {
	// Anything
	Interface any
}

// ClientReceiveCloudEventsResponse contains the response from method Client.ReceiveCloudEvents.
type ClientReceiveCloudEventsResponse struct {
	ReceiveResult
}

// ClientRejectCloudEventsResponse contains the response from method Client.RejectCloudEvents.
type ClientRejectCloudEventsResponse struct {
	RejectResult
}

// ClientReleaseCloudEventsResponse contains the response from method Client.ReleaseCloudEvents.
type ClientReleaseCloudEventsResponse struct {
	ReleaseResult
}
