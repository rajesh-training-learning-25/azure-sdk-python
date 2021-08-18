package appplatformapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/appplatform/mgmt/2021-06-01-preview/appplatform"
	"github.com/Azure/go-autorest/autorest"
)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, availabilityParameters appplatform.NameAvailabilityParameters) (result appplatform.NameAvailability, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, resource appplatform.ServiceResource) (result appplatform.ServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServicesDeleteFuture, err error)
	DisableTestEndpoint(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	EnableTestEndpoint(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.TestKeys, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServiceResource, err error)
	List(ctx context.Context, resourceGroupName string) (result appplatform.ServiceResourceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result appplatform.ServiceResourceListIterator, err error)
	ListBySubscription(ctx context.Context) (result appplatform.ServiceResourceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result appplatform.ServiceResourceListIterator, err error)
	ListTestKeys(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.TestKeys, err error)
	RegenerateTestKey(ctx context.Context, resourceGroupName string, serviceName string, regenerateTestKeyRequest appplatform.RegenerateTestKeyRequestPayload) (result appplatform.TestKeys, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, resource appplatform.ServiceResource) (result appplatform.ServicesUpdateFuture, err error)
}

var _ ServicesClientAPI = (*appplatform.ServicesClient)(nil)

// ConfigServersClientAPI contains the set of methods on the ConfigServersClient type.
type ConfigServersClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ConfigServerResource, err error)
	UpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, configServerResource appplatform.ConfigServerResource) (result appplatform.ConfigServersUpdatePatchFuture, err error)
	UpdatePut(ctx context.Context, resourceGroupName string, serviceName string, configServerResource appplatform.ConfigServerResource) (result appplatform.ConfigServersUpdatePutFuture, err error)
	Validate(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings appplatform.ConfigServerSettings) (result appplatform.ConfigServersValidateFuture, err error)
}

var _ ConfigServersClientAPI = (*appplatform.ConfigServersClient)(nil)

// MonitoringSettingsClientAPI contains the set of methods on the MonitoringSettingsClient type.
type MonitoringSettingsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.MonitoringSettingResource, err error)
	UpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource appplatform.MonitoringSettingResource) (result appplatform.MonitoringSettingsUpdatePatchFuture, err error)
	UpdatePut(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource appplatform.MonitoringSettingResource) (result appplatform.MonitoringSettingsUpdatePutFuture, err error)
}

var _ MonitoringSettingsClientAPI = (*appplatform.MonitoringSettingsClient)(nil)

// AppsClientAPI contains the set of methods on the AppsClient type.
type AppsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource appplatform.AppResource) (result appplatform.AppsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.AppsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, syncStatus string) (result appplatform.AppResource, err error)
	GetResourceUploadURL(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.ResourceUploadDefinition, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.AppResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.AppResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource appplatform.AppResource) (result appplatform.AppsUpdateFuture, err error)
	ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, appName string, validatePayload appplatform.CustomDomainValidatePayload) (result appplatform.CustomDomainValidateResult, err error)
}

var _ AppsClientAPI = (*appplatform.AppsClient)(nil)

// BindingsClientAPI contains the set of methods on the BindingsClient type.
type BindingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource appplatform.BindingResource) (result appplatform.BindingsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string) (result appplatform.BindingsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string) (result appplatform.BindingResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.BindingResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.BindingResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource appplatform.BindingResource) (result appplatform.BindingsUpdateFuture, err error)
}

var _ BindingsClientAPI = (*appplatform.BindingsClient)(nil)

// CertificatesClientAPI contains the set of methods on the CertificatesClient type.
type CertificatesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, certificateResource appplatform.CertificateResource) (result appplatform.CertificatesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, certificateName string) (result appplatform.CertificatesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, certificateName string) (result appplatform.CertificateResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.CertificateResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.CertificateResourceCollectionIterator, err error)
}

var _ CertificatesClientAPI = (*appplatform.CertificatesClient)(nil)

// CustomDomainsClientAPI contains the set of methods on the CustomDomainsClient type.
type CustomDomainsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource appplatform.CustomDomainResource) (result appplatform.CustomDomainsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string) (result appplatform.CustomDomainsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string) (result appplatform.CustomDomainResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.CustomDomainResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.CustomDomainResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource appplatform.CustomDomainResource) (result appplatform.CustomDomainsUpdateFuture, err error)
}

var _ CustomDomainsClientAPI = (*appplatform.CustomDomainsClient)(nil)

// DeploymentsClientAPI contains the set of methods on the DeploymentsClient type.
type DeploymentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource appplatform.DeploymentResource) (result appplatform.DeploymentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentResource, err error)
	GetLogFileURL(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.LogFileURLResponse, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, appName string, version []string) (result appplatform.DeploymentResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, appName string, version []string) (result appplatform.DeploymentResourceCollectionIterator, err error)
	ListForCluster(ctx context.Context, resourceGroupName string, serviceName string, version []string) (result appplatform.DeploymentResourceCollectionPage, err error)
	ListForClusterComplete(ctx context.Context, resourceGroupName string, serviceName string, version []string) (result appplatform.DeploymentResourceCollectionIterator, err error)
	Restart(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource appplatform.DeploymentResource) (result appplatform.DeploymentsUpdateFuture, err error)
}

var _ DeploymentsClientAPI = (*appplatform.DeploymentsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result appplatform.AvailableOperationsPage, err error)
	ListComplete(ctx context.Context) (result appplatform.AvailableOperationsIterator, err error)
}

var _ OperationsClientAPI = (*appplatform.OperationsClient)(nil)

// RuntimeVersionsClientAPI contains the set of methods on the RuntimeVersionsClient type.
type RuntimeVersionsClientAPI interface {
	ListRuntimeVersions(ctx context.Context) (result appplatform.AvailableRuntimeVersions, err error)
}

var _ RuntimeVersionsClientAPI = (*appplatform.RuntimeVersionsClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	List(ctx context.Context) (result appplatform.ResourceSkuCollectionPage, err error)
	ListComplete(ctx context.Context) (result appplatform.ResourceSkuCollectionIterator, err error)
}

var _ SkusClientAPI = (*appplatform.SkusClient)(nil)
