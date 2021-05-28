// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aztable

// GeoReplicationStatusType - The status of the secondary location.
type GeoReplicationStatusType string

const (
	GeoReplicationStatusTypeBootstrap   GeoReplicationStatusType = "bootstrap"
	GeoReplicationStatusTypeLive        GeoReplicationStatusType = "live"
	GeoReplicationStatusTypeUnavailable GeoReplicationStatusType = "unavailable"
)

// PossibleGeoReplicationStatusTypeValues returns the possible values for the GeoReplicationStatusType const type.
func PossibleGeoReplicationStatusTypeValues() []GeoReplicationStatusType {
	return []GeoReplicationStatusType{
		GeoReplicationStatusTypeBootstrap,
		GeoReplicationStatusTypeLive,
		GeoReplicationStatusTypeUnavailable,
	}
}

// ToPtr returns a *GeoReplicationStatusType pointing to the current value.
func (c GeoReplicationStatusType) ToPtr() *GeoReplicationStatusType {
	return &c
}

type OdataMetadataFormat string

const (
	OdataMetadataFormatApplicationJSONOdataFullmetadata    OdataMetadataFormat = "application/json;odata=fullmetadata"
	OdataMetadataFormatApplicationJSONOdataMinimalmetadata OdataMetadataFormat = "application/json;odata=minimalmetadata"
	OdataMetadataFormatApplicationJSONOdataNometadata      OdataMetadataFormat = "application/json;odata=nometadata"
)

// PossibleOdataMetadataFormatValues returns the possible values for the OdataMetadataFormat const type.
func PossibleOdataMetadataFormatValues() []OdataMetadataFormat {
	return []OdataMetadataFormat{
		OdataMetadataFormatApplicationJSONOdataFullmetadata,
		OdataMetadataFormatApplicationJSONOdataMinimalmetadata,
		OdataMetadataFormatApplicationJSONOdataNometadata,
	}
}

// ToPtr returns a *OdataMetadataFormat pointing to the current value.
func (c OdataMetadataFormat) ToPtr() *OdataMetadataFormat {
	return &c
}

type ResponseFormat string

const (
	ResponseFormatReturnContent   ResponseFormat = "return-content"
	ResponseFormatReturnNoContent ResponseFormat = "return-no-content"
)

// PossibleResponseFormatValues returns the possible values for the ResponseFormat const type.
func PossibleResponseFormatValues() []ResponseFormat {
	return []ResponseFormat{
		ResponseFormatReturnContent,
		ResponseFormatReturnNoContent,
	}
}

// ToPtr returns a *ResponseFormat pointing to the current value.
func (c ResponseFormat) ToPtr() *ResponseFormat {
	return &c
}
