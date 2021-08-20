// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package resourcegraph

import original "github.com/Azure/azure-sdk-for-go/services/resourcegraph/mgmt/2021-03-01/resourcegraph"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ColumnDataType = original.ColumnDataType

const (
	Boolean ColumnDataType = original.Boolean
	Integer ColumnDataType = original.Integer
	Number  ColumnDataType = original.Number
	Object  ColumnDataType = original.Object
	String  ColumnDataType = original.String
)

type FacetSortOrder = original.FacetSortOrder

const (
	Asc  FacetSortOrder = original.Asc
	Desc FacetSortOrder = original.Desc
)

type ResultFormat = original.ResultFormat

const (
	ResultFormatObjectArray ResultFormat = original.ResultFormatObjectArray
	ResultFormatTable       ResultFormat = original.ResultFormatTable
)

type ResultTruncated = original.ResultTruncated

const (
	False ResultTruncated = original.False
	True  ResultTruncated = original.True
)

type ResultType = original.ResultType

const (
	ResultTypeFacet       ResultType = original.ResultTypeFacet
	ResultTypeFacetError  ResultType = original.ResultTypeFacetError
	ResultTypeFacetResult ResultType = original.ResultTypeFacetResult
)

type BaseClient = original.BaseClient
type BasicFacet = original.BasicFacet
type Column = original.Column
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Facet = original.Facet
type FacetError = original.FacetError
type FacetRequest = original.FacetRequest
type FacetRequestOptions = original.FacetRequestOptions
type FacetResult = original.FacetResult
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type QueryRequest = original.QueryRequest
type QueryRequestOptions = original.QueryRequestOptions
type QueryResponse = original.QueryResponse
type Table = original.Table

func New() BaseClient {
	return original.New()
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleColumnDataTypeValues() []ColumnDataType {
	return original.PossibleColumnDataTypeValues()
}
func PossibleFacetSortOrderValues() []FacetSortOrder {
	return original.PossibleFacetSortOrderValues()
}
func PossibleResultFormatValues() []ResultFormat {
	return original.PossibleResultFormatValues()
}
func PossibleResultTruncatedValues() []ResultTruncated {
	return original.PossibleResultTruncatedValues()
}
func PossibleResultTypeValues() []ResultType {
	return original.PossibleResultTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
