//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azlogs

// BatchQueryRequest - An single request in a batch.
type BatchQueryRequest struct {
	// REQUIRED; The Analytics query. Learn more about the Analytics query syntax [https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/]
	Body *QueryBody

	// REQUIRED; Unique ID corresponding to each request in the batch.
	CorrelationID *string

	// REQUIRED; Primary Workspace ID of the query. This is the Workspace ID from the Properties blade in the Azure portal.
	WorkspaceID *string

	// Headers of the request. Can use prefer header to set server timeout and to query statistics and visualization information.
	Headers map[string]*string

	// The method of a single request in a batch, defaults to POST
	Method *BatchQueryRequestMethod

	// The query path of a single request in a batch, defaults to /query
	Path *BatchQueryRequestPath
}

// BatchQueryResponse - Contains the batch query response and the headers, id, and status of the request
type BatchQueryResponse struct {
	// Contains the tables, columns & rows resulting from a query.
	Body          *BatchQueryResults
	CorrelationID *string

	// Dictionary of
	Headers map[string]*string
	Status  *int32
}

// BatchQueryResults - Contains the tables, columns & rows resulting from a query.
type BatchQueryResults struct {
	// The code and message for an error.
	Error *ErrorInfo

	// Statistics represented in JSON format.
	Statistics []byte

	// The results of the query in tabular format.
	Tables []*Table

	// Visualization data in JSON format.
	Visualization []byte
}

// BatchRequest - An array of requests.
type BatchRequest struct {
	// REQUIRED; An single request in a batch.
	Requests []*BatchQueryRequest
}

// BatchResponse - Response to a batch query.
type BatchResponse struct {
	// An array of responses corresponding to each individual request in a batch.
	Responses []*BatchQueryResponse
}

// Column - A column in a table.
type Column struct {
	// The name of this column.
	Name *string

	// The data type of this column.
	Type *LogsColumnType
}

// QueryBody - The Analytics query. Learn more about the Analytics query syntax [https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/]
type QueryBody struct {
	// REQUIRED; The query to execute.
	Query *string

	// A list of workspaces to query in addition to the primary workspace.
	AdditionalWorkspaces []*string

	// Optional. The timespan over which to query data. This is an ISO8601 time period value. This timespan is applied in addition
	// to any that are specified in the query expression.
	Timespan *TimeInterval
}

// QueryResults - Contains the tables, columns & rows resulting from a query.
type QueryResults struct {
	// REQUIRED; The results of the query in tabular format.
	Tables []*Table

	// The code and message for an error.
	Error *ErrorInfo

	// Statistics represented in JSON format.
	Statistics []byte

	// Visualization data in JSON format.
	Visualization []byte
}

// Table - Contains the columns and rows for one table in a query response.
type Table struct {
	// REQUIRED; The list of columns in this table.
	Columns []*Column

	// REQUIRED; The name of the table.
	Name *string

	// REQUIRED; The resulting rows from this query.
	Rows []Row
}
