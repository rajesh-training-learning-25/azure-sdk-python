package virtualmachineimagebuilderapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/virtualmachineimagebuilder/mgmt/2021-10-01/virtualmachineimagebuilder"
)

// VirtualMachineImageTemplatesClientAPI contains the set of methods on the VirtualMachineImageTemplatesClient type.
type VirtualMachineImageTemplatesClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.VirtualMachineImageTemplatesCancelFuture, err error)
	CreateOrUpdate(ctx context.Context, parameters virtualmachineimagebuilder.ImageTemplate, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.VirtualMachineImageTemplatesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.VirtualMachineImageTemplatesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.ImageTemplate, err error)
	GetRunOutput(ctx context.Context, resourceGroupName string, imageTemplateName string, runOutputName string) (result virtualmachineimagebuilder.RunOutput, err error)
	List(ctx context.Context) (result virtualmachineimagebuilder.ImageTemplateListResultPage, err error)
	ListComplete(ctx context.Context) (result virtualmachineimagebuilder.ImageTemplateListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result virtualmachineimagebuilder.ImageTemplateListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result virtualmachineimagebuilder.ImageTemplateListResultIterator, err error)
	ListRunOutputs(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.RunOutputCollectionPage, err error)
	ListRunOutputsComplete(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.RunOutputCollectionIterator, err error)
	Run(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.VirtualMachineImageTemplatesRunFuture, err error)
	Update(ctx context.Context, parameters virtualmachineimagebuilder.ImageTemplateUpdateParameters, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.VirtualMachineImageTemplatesUpdateFuture, err error)
}

var _ VirtualMachineImageTemplatesClientAPI = (*virtualmachineimagebuilder.VirtualMachineImageTemplatesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result virtualmachineimagebuilder.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result virtualmachineimagebuilder.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*virtualmachineimagebuilder.OperationsClient)(nil)
