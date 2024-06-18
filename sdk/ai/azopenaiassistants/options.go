//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenaiassistants

// CancelRunOptions contains the optional parameters for the Client.CancelRun method.
type CancelRunOptions struct {
	// placeholder for future optional parameters
}

// CancelVectorStoreFileBatchOptions contains the optional parameters for the Client.CancelVectorStoreFileBatch method.
type CancelVectorStoreFileBatchOptions struct {
	// placeholder for future optional parameters
}

// CreateAssistantOptions contains the optional parameters for the Client.CreateAssistant method.
type CreateAssistantOptions struct {
	// placeholder for future optional parameters
}

// CreateMessageOptions contains the optional parameters for the Client.CreateMessage method.
type CreateMessageOptions struct {
	// placeholder for future optional parameters
}

// CreateRunOptions contains the optional parameters for the Client.CreateRun method.
type CreateRunOptions struct {
	// placeholder for future optional parameters
}

// CreateThreadAndRunOptions contains the optional parameters for the Client.CreateThreadAndRun method.
type CreateThreadAndRunOptions struct {
	// placeholder for future optional parameters
}

// CreateThreadOptions contains the optional parameters for the Client.CreateThread method.
type CreateThreadOptions struct {
	// placeholder for future optional parameters
}

// CreateVectorStoreFileBatchOptions contains the optional parameters for the Client.CreateVectorStoreFileBatch method.
type CreateVectorStoreFileBatchOptions struct {
	// placeholder for future optional parameters
}

// CreateVectorStoreFileOptions contains the optional parameters for the Client.CreateVectorStoreFile method.
type CreateVectorStoreFileOptions struct {
	// placeholder for future optional parameters
}

// CreateVectorStoreOptions contains the optional parameters for the Client.CreateVectorStore method.
type CreateVectorStoreOptions struct {
	// placeholder for future optional parameters
}

// DeleteAssistantOptions contains the optional parameters for the Client.DeleteAssistant method.
type DeleteAssistantOptions struct {
	// placeholder for future optional parameters
}

// DeleteFileOptions contains the optional parameters for the Client.DeleteFile method.
type DeleteFileOptions struct {
	// placeholder for future optional parameters
}

// DeleteThreadOptions contains the optional parameters for the Client.DeleteThread method.
type DeleteThreadOptions struct {
	// placeholder for future optional parameters
}

// DeleteVectorStoreFileOptions contains the optional parameters for the Client.DeleteVectorStoreFile method.
type DeleteVectorStoreFileOptions struct {
	// placeholder for future optional parameters
}

// DeleteVectorStoreOptions contains the optional parameters for the Client.DeleteVectorStore method.
type DeleteVectorStoreOptions struct {
	// placeholder for future optional parameters
}

// GetAssistantOptions contains the optional parameters for the Client.GetAssistant method.
type GetAssistantOptions struct {
	// placeholder for future optional parameters
}

// GetFileOptions contains the optional parameters for the Client.GetFile method.
type GetFileOptions struct {
	// placeholder for future optional parameters
}

// GetMessageOptions contains the optional parameters for the Client.GetMessage method.
type GetMessageOptions struct {
	// placeholder for future optional parameters
}

// GetRunOptions contains the optional parameters for the Client.GetRun method.
type GetRunOptions struct {
	// placeholder for future optional parameters
}

// GetRunStepOptions contains the optional parameters for the Client.GetRunStep method.
type GetRunStepOptions struct {
	// placeholder for future optional parameters
}

// GetThreadOptions contains the optional parameters for the Client.GetThread method.
type GetThreadOptions struct {
	// placeholder for future optional parameters
}

// GetVectorStoreFileBatchOptions contains the optional parameters for the Client.GetVectorStoreFileBatch method.
type GetVectorStoreFileBatchOptions struct {
	// placeholder for future optional parameters
}

// GetVectorStoreFileOptions contains the optional parameters for the Client.GetVectorStoreFile method.
type GetVectorStoreFileOptions struct {
	// placeholder for future optional parameters
}

// GetVectorStoreOptions contains the optional parameters for the Client.GetVectorStore method.
type GetVectorStoreOptions struct {
	// placeholder for future optional parameters
}

// ListAssistantsOptions contains the optional parameters for the Client.ListAssistants method.
type ListAssistantsOptions struct {
	// A cursor for use in pagination. after is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include after=objfoo in order to fetch the next page of the list.
	After *string

	// A cursor for use in pagination. before is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include before=objfoo in order to fetch the previous page of the list.
	Before *string

	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	Limit *int32

	// Sort order by the created_at timestamp of the objects. asc for ascending order and desc for descending order.
	Order *ListSortOrder
}

// ListFilesOptions contains the optional parameters for the Client.ListFiles method.
type ListFilesOptions struct {
	// A value that, when provided, limits list results to files matching the corresponding purpose.
	Purpose *FilePurpose
}

// ListMessagesOptions contains the optional parameters for the Client.ListMessages method.
type ListMessagesOptions struct {
	// A cursor for use in pagination. after is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include after=objfoo in order to fetch the next page of the list.
	After *string

	// A cursor for use in pagination. before is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include before=objfoo in order to fetch the previous page of the list.
	Before *string

	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	Limit *int32

	// Sort order by the created_at timestamp of the objects. asc for ascending order and desc for descending order.
	Order *ListSortOrder

	// Filter messages by the run ID that generated them.
	RunID *string
}

// ListRunStepsOptions contains the optional parameters for the Client.ListRunSteps method.
type ListRunStepsOptions struct {
	// A cursor for use in pagination. after is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include after=objfoo in order to fetch the next page of the list.
	After *string

	// A cursor for use in pagination. before is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include before=objfoo in order to fetch the previous page of the list.
	Before *string

	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	Limit *int32

	// Sort order by the created_at timestamp of the objects. asc for ascending order and desc for descending order.
	Order *ListSortOrder
}

// ListRunsOptions contains the optional parameters for the Client.ListRuns method.
type ListRunsOptions struct {
	// A cursor for use in pagination. after is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include after=objfoo in order to fetch the next page of the list.
	After *string

	// A cursor for use in pagination. before is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include before=objfoo in order to fetch the previous page of the list.
	Before *string

	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	Limit *int32

	// Sort order by the created_at timestamp of the objects. asc for ascending order and desc for descending order.
	Order *ListSortOrder
}

// ListVectorStoreFileBatchFilesOptions contains the optional parameters for the Client.ListVectorStoreFileBatchFiles
// method.
type ListVectorStoreFileBatchFilesOptions struct {
	// A cursor for use in pagination. after is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include after=objfoo in order to fetch the next page of the list.
	After *string

	// A cursor for use in pagination. before is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include before=objfoo in order to fetch the previous page of the list.
	Before *string

	// Filter by file status.
	Filter *VectorStoreFileStatusFilter

	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	Limit *int32

	// Sort order by the created_at timestamp of the objects. asc for ascending order and desc for descending order.
	Order *ListSortOrder
}

// ListVectorStoreFilesOptions contains the optional parameters for the Client.ListVectorStoreFiles method.
type ListVectorStoreFilesOptions struct {
	// A cursor for use in pagination. after is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include after=objfoo in order to fetch the next page of the list.
	After *string

	// A cursor for use in pagination. before is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include before=objfoo in order to fetch the previous page of the list.
	Before *string

	// Filter by file status.
	Filter *VectorStoreFileStatusFilter

	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	Limit *int32

	// Sort order by the created_at timestamp of the objects. asc for ascending order and desc for descending order.
	Order *ListSortOrder
}

// ListVectorStoresOptions contains the optional parameters for the Client.ListVectorStores method.
type ListVectorStoresOptions struct {
	// A cursor for use in pagination. after is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include after=objfoo in order to fetch the next page of the list.
	After *string

	// A cursor for use in pagination. before is an object ID that defines your place in the list. For instance, if you make a
	// list request and receive 100 objects, ending with objfoo, your subsequent call
	// can include before=objfoo in order to fetch the previous page of the list.
	Before *string

	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	Limit *int32

	// Sort order by the created_at timestamp of the objects. asc for ascending order and desc for descending order.
	Order *ListSortOrder
}

// ModifyVectorStoreOptions contains the optional parameters for the Client.ModifyVectorStore method.
type ModifyVectorStoreOptions struct {
	// placeholder for future optional parameters
}

// SubmitToolOutputsToRunOptions contains the optional parameters for the Client.SubmitToolOutputsToRun method.
type SubmitToolOutputsToRunOptions struct {
	// placeholder for future optional parameters
}

// UpdateAssistantOptions contains the optional parameters for the Client.UpdateAssistant method.
type UpdateAssistantOptions struct {
	// placeholder for future optional parameters
}

// UpdateMessageOptions contains the optional parameters for the Client.UpdateMessage method.
type UpdateMessageOptions struct {
	// placeholder for future optional parameters
}

// UpdateRunOptions contains the optional parameters for the Client.UpdateRun method.
type UpdateRunOptions struct {
	// placeholder for future optional parameters
}

// UpdateThreadOptions contains the optional parameters for the Client.UpdateThread method.
type UpdateThreadOptions struct {
	// placeholder for future optional parameters
}

// UploadFileOptions contains the optional parameters for the Client.UploadFile method.
type UploadFileOptions struct {
	// A filename to associate with the uploaded data.
	Filename *string
}
