//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package storsimple

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/storsimple1200series/mgmt/2016-10-01/storsimple"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AlertEmailNotificationStatus = original.AlertEmailNotificationStatus

const (
	Disabled AlertEmailNotificationStatus = original.Disabled
	Enabled  AlertEmailNotificationStatus = original.Enabled
)

type AlertScope = original.AlertScope

const (
	AlertScopeDevice   AlertScope = original.AlertScopeDevice
	AlertScopeResource AlertScope = original.AlertScopeResource
)

type AlertSeverity = original.AlertSeverity

const (
	Critical      AlertSeverity = original.Critical
	Informational AlertSeverity = original.Informational
	Warning       AlertSeverity = original.Warning
)

type AlertSourceType = original.AlertSourceType

const (
	AlertSourceTypeDevice   AlertSourceType = original.AlertSourceTypeDevice
	AlertSourceTypeResource AlertSourceType = original.AlertSourceTypeResource
)

type AlertStatus = original.AlertStatus

const (
	Active  AlertStatus = original.Active
	Cleared AlertStatus = original.Cleared
)

type AuthType = original.AuthType

const (
	AccessControlService AuthType = original.AccessControlService
	AzureActiveDirectory AuthType = original.AzureActiveDirectory
	Invalid              AuthType = original.Invalid
)

type CloudType = original.CloudType

const (
	Azure     CloudType = original.Azure
	HP        CloudType = original.HP
	OpenStack CloudType = original.OpenStack
	S3        CloudType = original.S3
	S3RRS     CloudType = original.S3RRS
)

type ContractVersions = original.ContractVersions

const (
	InvalidVersion ContractVersions = original.InvalidVersion
	V201109        ContractVersions = original.V201109
	V201202        ContractVersions = original.V201202
	V201205        ContractVersions = original.V201205
	V201212        ContractVersions = original.V201212
	V201304        ContractVersions = original.V201304
	V201310        ContractVersions = original.V201310
	V201311        ContractVersions = original.V201311
	V201404        ContractVersions = original.V201404
	V201406        ContractVersions = original.V201406
	V201407        ContractVersions = original.V201407
	V201409        ContractVersions = original.V201409
	V201410        ContractVersions = original.V201410
	V201412        ContractVersions = original.V201412
	V201501        ContractVersions = original.V201501
	V201502        ContractVersions = original.V201502
	V201504        ContractVersions = original.V201504
	V201505        ContractVersions = original.V201505
	V201506        ContractVersions = original.V201506
	V201507        ContractVersions = original.V201507
	V201508        ContractVersions = original.V201508
	V201510        ContractVersions = original.V201510
	V201512        ContractVersions = original.V201512
	V201601        ContractVersions = original.V201601
	V201602        ContractVersions = original.V201602
	V201604        ContractVersions = original.V201604
	V201605        ContractVersions = original.V201605
	V201607        ContractVersions = original.V201607
	V201608        ContractVersions = original.V201608
)

type DataPolicy = original.DataPolicy

const (
	DataPolicyCloud   DataPolicy = original.DataPolicyCloud
	DataPolicyInvalid DataPolicy = original.DataPolicyInvalid
	DataPolicyLocal   DataPolicy = original.DataPolicyLocal
	DataPolicyTiered  DataPolicy = original.DataPolicyTiered
)

type DeviceConfigurationStatus = original.DeviceConfigurationStatus

const (
	Complete DeviceConfigurationStatus = original.Complete
	Pending  DeviceConfigurationStatus = original.Pending
)

type DeviceOperation = original.DeviceOperation

const (
	Browsable         DeviceOperation = original.Browsable
	Deactivate        DeviceOperation = original.Deactivate
	Delete            DeviceOperation = original.Delete
	DeleteWithWarning DeviceOperation = original.DeleteWithWarning
	DRSource          DeviceOperation = original.DRSource
	DRTarget          DeviceOperation = original.DRTarget
	None              DeviceOperation = original.None
	ReadOnlyForDR     DeviceOperation = original.ReadOnlyForDR
)

type DeviceStatus = original.DeviceStatus

const (
	Creating          DeviceStatus = original.Creating
	Deactivated       DeviceStatus = original.Deactivated
	Deactivating      DeviceStatus = original.Deactivating
	Deleted           DeviceStatus = original.Deleted
	MaintenanceMode   DeviceStatus = original.MaintenanceMode
	Offline           DeviceStatus = original.Offline
	Online            DeviceStatus = original.Online
	Provisioning      DeviceStatus = original.Provisioning
	ReadyToSetup      DeviceStatus = original.ReadyToSetup
	RequiresAttention DeviceStatus = original.RequiresAttention
	Unknown           DeviceStatus = original.Unknown
)

type DeviceType = original.DeviceType

const (
	DeviceTypeAppliance                        DeviceType = original.DeviceTypeAppliance
	DeviceTypeInvalid                          DeviceType = original.DeviceTypeInvalid
	DeviceTypeSeries9000OnPremVirtualAppliance DeviceType = original.DeviceTypeSeries9000OnPremVirtualAppliance
	DeviceTypeSeries9000PhysicalAppliance      DeviceType = original.DeviceTypeSeries9000PhysicalAppliance
	DeviceTypeSeries9000VirtualAppliance       DeviceType = original.DeviceTypeSeries9000VirtualAppliance
	DeviceTypeVirtualAppliance                 DeviceType = original.DeviceTypeVirtualAppliance
)

type DhcpStatus = original.DhcpStatus

const (
	DhcpStatusDisabled DhcpStatus = original.DhcpStatusDisabled
	DhcpStatusEnabled  DhcpStatus = original.DhcpStatusEnabled
)

type DiskStatus = original.DiskStatus

const (
	DiskStatusOffline DiskStatus = original.DiskStatusOffline
	DiskStatusOnline  DiskStatus = original.DiskStatusOnline
)

type DownloadPhase = original.DownloadPhase

const (
	DownloadPhaseDownloading  DownloadPhase = original.DownloadPhaseDownloading
	DownloadPhaseInitializing DownloadPhase = original.DownloadPhaseInitializing
	DownloadPhaseUnknown      DownloadPhase = original.DownloadPhaseUnknown
	DownloadPhaseVerifying    DownloadPhase = original.DownloadPhaseVerifying
)

type EncryptionAlgorithm = original.EncryptionAlgorithm

const (
	EncryptionAlgorithmAES256        EncryptionAlgorithm = original.EncryptionAlgorithmAES256
	EncryptionAlgorithmNone          EncryptionAlgorithm = original.EncryptionAlgorithmNone
	EncryptionAlgorithmRSAESPKCS1V15 EncryptionAlgorithm = original.EncryptionAlgorithmRSAESPKCS1V15
)

type EncryptionStatus = original.EncryptionStatus

const (
	EncryptionStatusDisabled EncryptionStatus = original.EncryptionStatusDisabled
	EncryptionStatusEnabled  EncryptionStatus = original.EncryptionStatusEnabled
)

type InitiatedBy = original.InitiatedBy

const (
	Manual    InitiatedBy = original.Manual
	Scheduled InitiatedBy = original.Scheduled
)

type JobStatus = original.JobStatus

const (
	JobStatusCanceled  JobStatus = original.JobStatusCanceled
	JobStatusFailed    JobStatus = original.JobStatusFailed
	JobStatusInvalid   JobStatus = original.JobStatusInvalid
	JobStatusPaused    JobStatus = original.JobStatusPaused
	JobStatusRunning   JobStatus = original.JobStatusRunning
	JobStatusScheduled JobStatus = original.JobStatusScheduled
	JobStatusSucceeded JobStatus = original.JobStatusSucceeded
)

type JobType = original.JobType

const (
	JobTypeBackup          JobType = original.JobTypeBackup
	JobTypeClone           JobType = original.JobTypeClone
	JobTypeDownloadUpdates JobType = original.JobTypeDownloadUpdates
	JobTypeFailover        JobType = original.JobTypeFailover
	JobTypeInstallUpdates  JobType = original.JobTypeInstallUpdates
)

type KeyRolloverStatus = original.KeyRolloverStatus

const (
	NotRequired KeyRolloverStatus = original.NotRequired
	Required    KeyRolloverStatus = original.Required
)

type ManagerType = original.ManagerType

const (
	GardaV1    ManagerType = original.GardaV1
	HelsinkiV1 ManagerType = original.HelsinkiV1
)

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage MetricAggregationType = original.MetricAggregationTypeAverage
	MetricAggregationTypeLast    MetricAggregationType = original.MetricAggregationTypeLast
	MetricAggregationTypeMaximum MetricAggregationType = original.MetricAggregationTypeMaximum
	MetricAggregationTypeMinimum MetricAggregationType = original.MetricAggregationTypeMinimum
	MetricAggregationTypeNone    MetricAggregationType = original.MetricAggregationTypeNone
	MetricAggregationTypeTotal   MetricAggregationType = original.MetricAggregationTypeTotal
)

type MetricUnit = original.MetricUnit

const (
	Bytes          MetricUnit = original.Bytes
	BytesPerSecond MetricUnit = original.BytesPerSecond
	Count          MetricUnit = original.Count
	CountPerSecond MetricUnit = original.CountPerSecond
	Percent        MetricUnit = original.Percent
	Seconds        MetricUnit = original.Seconds
)

type MonitoringStatus = original.MonitoringStatus

const (
	MonitoringStatusDisabled MonitoringStatus = original.MonitoringStatusDisabled
	MonitoringStatusEnabled  MonitoringStatus = original.MonitoringStatusEnabled
)

type ServiceOwnersAlertNotificationStatus = original.ServiceOwnersAlertNotificationStatus

const (
	ServiceOwnersAlertNotificationStatusDisabled ServiceOwnersAlertNotificationStatus = original.ServiceOwnersAlertNotificationStatusDisabled
	ServiceOwnersAlertNotificationStatusEnabled  ServiceOwnersAlertNotificationStatus = original.ServiceOwnersAlertNotificationStatusEnabled
)

type ShareStatus = original.ShareStatus

const (
	ShareStatusOffline ShareStatus = original.ShareStatusOffline
	ShareStatusOnline  ShareStatus = original.ShareStatusOnline
)

type SslStatus = original.SslStatus

const (
	SslStatusDisabled SslStatus = original.SslStatusDisabled
	SslStatusEnabled  SslStatus = original.SslStatusEnabled
)

type SupportedDeviceCapabilities = original.SupportedDeviceCapabilities

const (
	SupportedDeviceCapabilitiesFileServer  SupportedDeviceCapabilities = original.SupportedDeviceCapabilitiesFileServer
	SupportedDeviceCapabilitiesInvalid     SupportedDeviceCapabilities = original.SupportedDeviceCapabilitiesInvalid
	SupportedDeviceCapabilitiesIscsiServer SupportedDeviceCapabilities = original.SupportedDeviceCapabilitiesIscsiServer
)

type TargetType = original.TargetType

const (
	TargetTypeDiskServer TargetType = original.TargetTypeDiskServer
	TargetTypeFileServer TargetType = original.TargetTypeFileServer
)

type UpdateOperation = original.UpdateOperation

const (
	Downloading UpdateOperation = original.Downloading
	Idle        UpdateOperation = original.Idle
	Installing  UpdateOperation = original.Installing
	Scanning    UpdateOperation = original.Scanning
)

type AccessControlRecord = original.AccessControlRecord
type AccessControlRecordList = original.AccessControlRecordList
type AccessControlRecordProperties = original.AccessControlRecordProperties
type AccessControlRecordsClient = original.AccessControlRecordsClient
type AccessControlRecordsCreateOrUpdateFuture = original.AccessControlRecordsCreateOrUpdateFuture
type AccessControlRecordsDeleteFuture = original.AccessControlRecordsDeleteFuture
type Alert = original.Alert
type AlertErrorDetails = original.AlertErrorDetails
type AlertFilter = original.AlertFilter
type AlertList = original.AlertList
type AlertListIterator = original.AlertListIterator
type AlertListPage = original.AlertListPage
type AlertProperties = original.AlertProperties
type AlertSettings = original.AlertSettings
type AlertSettingsProperties = original.AlertSettingsProperties
type AlertSource = original.AlertSource
type AlertsClient = original.AlertsClient
type AsymmetricEncryptedSecret = original.AsymmetricEncryptedSecret
type AvailableProviderOperation = original.AvailableProviderOperation
type AvailableProviderOperationDisplay = original.AvailableProviderOperationDisplay
type AvailableProviderOperations = original.AvailableProviderOperations
type AvailableProviderOperationsClient = original.AvailableProviderOperationsClient
type AvailableProviderOperationsIterator = original.AvailableProviderOperationsIterator
type AvailableProviderOperationsPage = original.AvailableProviderOperationsPage
type Backup = original.Backup
type BackupElement = original.BackupElement
type BackupElementProperties = original.BackupElementProperties
type BackupFilter = original.BackupFilter
type BackupList = original.BackupList
type BackupListIterator = original.BackupListIterator
type BackupListPage = original.BackupListPage
type BackupProperties = original.BackupProperties
type BackupScheduleGroup = original.BackupScheduleGroup
type BackupScheduleGroupList = original.BackupScheduleGroupList
type BackupScheduleGroupProperties = original.BackupScheduleGroupProperties
type BackupScheduleGroupsClient = original.BackupScheduleGroupsClient
type BackupScheduleGroupsCreateOrUpdateFuture = original.BackupScheduleGroupsCreateOrUpdateFuture
type BackupScheduleGroupsDeleteFuture = original.BackupScheduleGroupsDeleteFuture
type BackupsClient = original.BackupsClient
type BackupsCloneFuture = original.BackupsCloneFuture
type BackupsDeleteFuture = original.BackupsDeleteFuture
type BaseClient = original.BaseClient
type BaseModel = original.BaseModel
type ChapProperties = original.ChapProperties
type ChapSettings = original.ChapSettings
type ChapSettingsClient = original.ChapSettingsClient
type ChapSettingsCreateOrUpdateFuture = original.ChapSettingsCreateOrUpdateFuture
type ChapSettingsDeleteFuture = original.ChapSettingsDeleteFuture
type ChapSettingsList = original.ChapSettingsList
type ClearAlertRequest = original.ClearAlertRequest
type CloneRequest = original.CloneRequest
type CloneRequestProperties = original.CloneRequestProperties
type Device = original.Device
type DeviceDetails = original.DeviceDetails
type DeviceList = original.DeviceList
type DevicePatch = original.DevicePatch
type DeviceProperties = original.DeviceProperties
type DevicesClient = original.DevicesClient
type DevicesCreateOrUpdateAlertSettingsFuture = original.DevicesCreateOrUpdateAlertSettingsFuture
type DevicesCreateOrUpdateSecuritySettingsFuture = original.DevicesCreateOrUpdateSecuritySettingsFuture
type DevicesDeactivateFuture = original.DevicesDeactivateFuture
type DevicesDeleteFuture = original.DevicesDeleteFuture
type DevicesDownloadUpdatesFuture = original.DevicesDownloadUpdatesFuture
type DevicesFailoverFuture = original.DevicesFailoverFuture
type DevicesInstallUpdatesFuture = original.DevicesInstallUpdatesFuture
type DevicesPatchFuture = original.DevicesPatchFuture
type DevicesScanForUpdatesFuture = original.DevicesScanForUpdatesFuture
type EncryptionSettings = original.EncryptionSettings
type EncryptionSettingsProperties = original.EncryptionSettingsProperties
type Error = original.Error
type FailoverRequest = original.FailoverRequest
type FileServer = original.FileServer
type FileServerList = original.FileServerList
type FileServerProperties = original.FileServerProperties
type FileServersBackupNowFuture = original.FileServersBackupNowFuture
type FileServersClient = original.FileServersClient
type FileServersCreateOrUpdateFuture = original.FileServersCreateOrUpdateFuture
type FileServersDeleteFuture = original.FileServersDeleteFuture
type FileShare = original.FileShare
type FileShareList = original.FileShareList
type FileShareProperties = original.FileShareProperties
type FileSharesClient = original.FileSharesClient
type FileSharesCreateOrUpdateFuture = original.FileSharesCreateOrUpdateFuture
type FileSharesDeleteFuture = original.FileSharesDeleteFuture
type IPConfig = original.IPConfig
type ISCSIDisk = original.ISCSIDisk
type ISCSIDiskList = original.ISCSIDiskList
type ISCSIDiskProperties = original.ISCSIDiskProperties
type ISCSIServer = original.ISCSIServer
type ISCSIServerList = original.ISCSIServerList
type ISCSIServerProperties = original.ISCSIServerProperties
type IscsiDisksClient = original.IscsiDisksClient
type IscsiDisksCreateOrUpdateFuture = original.IscsiDisksCreateOrUpdateFuture
type IscsiDisksDeleteFuture = original.IscsiDisksDeleteFuture
type IscsiServersBackupNowFuture = original.IscsiServersBackupNowFuture
type IscsiServersClient = original.IscsiServersClient
type IscsiServersCreateOrUpdateFuture = original.IscsiServersCreateOrUpdateFuture
type IscsiServersDeleteFuture = original.IscsiServersDeleteFuture
type Item = original.Item
type Job = original.Job
type JobErrorDetails = original.JobErrorDetails
type JobErrorItem = original.JobErrorItem
type JobFilter = original.JobFilter
type JobList = original.JobList
type JobListIterator = original.JobListIterator
type JobListPage = original.JobListPage
type JobProperties = original.JobProperties
type JobStage = original.JobStage
type JobStats = original.JobStats
type JobsClient = original.JobsClient
type Manager = original.Manager
type ManagerExtendedInfo = original.ManagerExtendedInfo
type ManagerExtendedInfoProperties = original.ManagerExtendedInfoProperties
type ManagerIntrinsicSettings = original.ManagerIntrinsicSettings
type ManagerList = original.ManagerList
type ManagerPatch = original.ManagerPatch
type ManagerProperties = original.ManagerProperties
type ManagerSku = original.ManagerSku
type ManagersClient = original.ManagersClient
type Message = original.Message
type MetricAvailablity = original.MetricAvailablity
type MetricData = original.MetricData
type MetricDefinition = original.MetricDefinition
type MetricDefinitionList = original.MetricDefinitionList
type MetricDimension = original.MetricDimension
type MetricFilter = original.MetricFilter
type MetricList = original.MetricList
type MetricName = original.MetricName
type MetricNameFilter = original.MetricNameFilter
type Metrics = original.Metrics
type NetworkAdapter = original.NetworkAdapter
type NetworkSettings = original.NetworkSettings
type NetworkSettingsProperties = original.NetworkSettingsProperties
type NodeNetwork = original.NodeNetwork
type RawCertificateData = original.RawCertificateData
type Resource = original.Resource
type ResourceCertificateAndAADDetails = original.ResourceCertificateAndAADDetails
type SecuritySettings = original.SecuritySettings
type SecuritySettingsProperties = original.SecuritySettingsProperties
type SendTestAlertEmailRequest = original.SendTestAlertEmailRequest
type StorageAccountCredential = original.StorageAccountCredential
type StorageAccountCredentialList = original.StorageAccountCredentialList
type StorageAccountCredentialProperties = original.StorageAccountCredentialProperties
type StorageAccountCredentialsClient = original.StorageAccountCredentialsClient
type StorageAccountCredentialsCreateOrUpdateFuture = original.StorageAccountCredentialsCreateOrUpdateFuture
type StorageAccountCredentialsDeleteFuture = original.StorageAccountCredentialsDeleteFuture
type StorageDomain = original.StorageDomain
type StorageDomainList = original.StorageDomainList
type StorageDomainProperties = original.StorageDomainProperties
type StorageDomainsClient = original.StorageDomainsClient
type StorageDomainsCreateOrUpdateFuture = original.StorageDomainsCreateOrUpdateFuture
type StorageDomainsDeleteFuture = original.StorageDomainsDeleteFuture
type SymmetricEncryptedSecret = original.SymmetricEncryptedSecret
type Time = original.Time
type TimeSettings = original.TimeSettings
type TimeSettingsProperties = original.TimeSettingsProperties
type UpdateDownloadProgress = original.UpdateDownloadProgress
type UpdateInstallProgress = original.UpdateInstallProgress
type Updates = original.Updates
type UpdatesProperties = original.UpdatesProperties
type UploadCertificateRequest = original.UploadCertificateRequest
type UploadCertificateResponse = original.UploadCertificateResponse

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccessControlRecordsClient(subscriptionID string) AccessControlRecordsClient {
	return original.NewAccessControlRecordsClient(subscriptionID)
}
func NewAccessControlRecordsClientWithBaseURI(baseURI string, subscriptionID string) AccessControlRecordsClient {
	return original.NewAccessControlRecordsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertListIterator(page AlertListPage) AlertListIterator {
	return original.NewAlertListIterator(page)
}
func NewAlertListPage(cur AlertList, getNextPage func(context.Context, AlertList) (AlertList, error)) AlertListPage {
	return original.NewAlertListPage(cur, getNextPage)
}
func NewAlertsClient(subscriptionID string) AlertsClient {
	return original.NewAlertsClient(subscriptionID)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAvailableProviderOperationsClient(subscriptionID string) AvailableProviderOperationsClient {
	return original.NewAvailableProviderOperationsClient(subscriptionID)
}
func NewAvailableProviderOperationsClientWithBaseURI(baseURI string, subscriptionID string) AvailableProviderOperationsClient {
	return original.NewAvailableProviderOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAvailableProviderOperationsIterator(page AvailableProviderOperationsPage) AvailableProviderOperationsIterator {
	return original.NewAvailableProviderOperationsIterator(page)
}
func NewAvailableProviderOperationsPage(cur AvailableProviderOperations, getNextPage func(context.Context, AvailableProviderOperations) (AvailableProviderOperations, error)) AvailableProviderOperationsPage {
	return original.NewAvailableProviderOperationsPage(cur, getNextPage)
}
func NewBackupListIterator(page BackupListPage) BackupListIterator {
	return original.NewBackupListIterator(page)
}
func NewBackupListPage(cur BackupList, getNextPage func(context.Context, BackupList) (BackupList, error)) BackupListPage {
	return original.NewBackupListPage(cur, getNextPage)
}
func NewBackupScheduleGroupsClient(subscriptionID string) BackupScheduleGroupsClient {
	return original.NewBackupScheduleGroupsClient(subscriptionID)
}
func NewBackupScheduleGroupsClientWithBaseURI(baseURI string, subscriptionID string) BackupScheduleGroupsClient {
	return original.NewBackupScheduleGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBackupsClient(subscriptionID string) BackupsClient {
	return original.NewBackupsClient(subscriptionID)
}
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return original.NewBackupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewChapSettingsClient(subscriptionID string) ChapSettingsClient {
	return original.NewChapSettingsClient(subscriptionID)
}
func NewChapSettingsClientWithBaseURI(baseURI string, subscriptionID string) ChapSettingsClient {
	return original.NewChapSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDevicesClient(subscriptionID string) DevicesClient {
	return original.NewDevicesClient(subscriptionID)
}
func NewDevicesClientWithBaseURI(baseURI string, subscriptionID string) DevicesClient {
	return original.NewDevicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFileServersClient(subscriptionID string) FileServersClient {
	return original.NewFileServersClient(subscriptionID)
}
func NewFileServersClientWithBaseURI(baseURI string, subscriptionID string) FileServersClient {
	return original.NewFileServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewFileSharesClient(subscriptionID string) FileSharesClient {
	return original.NewFileSharesClient(subscriptionID)
}
func NewFileSharesClientWithBaseURI(baseURI string, subscriptionID string) FileSharesClient {
	return original.NewFileSharesClientWithBaseURI(baseURI, subscriptionID)
}
func NewIscsiDisksClient(subscriptionID string) IscsiDisksClient {
	return original.NewIscsiDisksClient(subscriptionID)
}
func NewIscsiDisksClientWithBaseURI(baseURI string, subscriptionID string) IscsiDisksClient {
	return original.NewIscsiDisksClientWithBaseURI(baseURI, subscriptionID)
}
func NewIscsiServersClient(subscriptionID string) IscsiServersClient {
	return original.NewIscsiServersClient(subscriptionID)
}
func NewIscsiServersClientWithBaseURI(baseURI string, subscriptionID string) IscsiServersClient {
	return original.NewIscsiServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobListIterator(page JobListPage) JobListIterator {
	return original.NewJobListIterator(page)
}
func NewJobListPage(cur JobList, getNextPage func(context.Context, JobList) (JobList, error)) JobListPage {
	return original.NewJobListPage(cur, getNextPage)
}
func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagersClient(subscriptionID string) ManagersClient {
	return original.NewManagersClient(subscriptionID)
}
func NewManagersClientWithBaseURI(baseURI string, subscriptionID string) ManagersClient {
	return original.NewManagersClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageAccountCredentialsClient(subscriptionID string) StorageAccountCredentialsClient {
	return original.NewStorageAccountCredentialsClient(subscriptionID)
}
func NewStorageAccountCredentialsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountCredentialsClient {
	return original.NewStorageAccountCredentialsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageDomainsClient(subscriptionID string) StorageDomainsClient {
	return original.NewStorageDomainsClient(subscriptionID)
}
func NewStorageDomainsClientWithBaseURI(baseURI string, subscriptionID string) StorageDomainsClient {
	return original.NewStorageDomainsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAlertEmailNotificationStatusValues() []AlertEmailNotificationStatus {
	return original.PossibleAlertEmailNotificationStatusValues()
}
func PossibleAlertScopeValues() []AlertScope {
	return original.PossibleAlertScopeValues()
}
func PossibleAlertSeverityValues() []AlertSeverity {
	return original.PossibleAlertSeverityValues()
}
func PossibleAlertSourceTypeValues() []AlertSourceType {
	return original.PossibleAlertSourceTypeValues()
}
func PossibleAlertStatusValues() []AlertStatus {
	return original.PossibleAlertStatusValues()
}
func PossibleAuthTypeValues() []AuthType {
	return original.PossibleAuthTypeValues()
}
func PossibleCloudTypeValues() []CloudType {
	return original.PossibleCloudTypeValues()
}
func PossibleContractVersionsValues() []ContractVersions {
	return original.PossibleContractVersionsValues()
}
func PossibleDataPolicyValues() []DataPolicy {
	return original.PossibleDataPolicyValues()
}
func PossibleDeviceConfigurationStatusValues() []DeviceConfigurationStatus {
	return original.PossibleDeviceConfigurationStatusValues()
}
func PossibleDeviceOperationValues() []DeviceOperation {
	return original.PossibleDeviceOperationValues()
}
func PossibleDeviceStatusValues() []DeviceStatus {
	return original.PossibleDeviceStatusValues()
}
func PossibleDeviceTypeValues() []DeviceType {
	return original.PossibleDeviceTypeValues()
}
func PossibleDhcpStatusValues() []DhcpStatus {
	return original.PossibleDhcpStatusValues()
}
func PossibleDiskStatusValues() []DiskStatus {
	return original.PossibleDiskStatusValues()
}
func PossibleDownloadPhaseValues() []DownloadPhase {
	return original.PossibleDownloadPhaseValues()
}
func PossibleEncryptionAlgorithmValues() []EncryptionAlgorithm {
	return original.PossibleEncryptionAlgorithmValues()
}
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return original.PossibleEncryptionStatusValues()
}
func PossibleInitiatedByValues() []InitiatedBy {
	return original.PossibleInitiatedByValues()
}
func PossibleJobStatusValues() []JobStatus {
	return original.PossibleJobStatusValues()
}
func PossibleJobTypeValues() []JobType {
	return original.PossibleJobTypeValues()
}
func PossibleKeyRolloverStatusValues() []KeyRolloverStatus {
	return original.PossibleKeyRolloverStatusValues()
}
func PossibleManagerTypeValues() []ManagerType {
	return original.PossibleManagerTypeValues()
}
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}
func PossibleMetricUnitValues() []MetricUnit {
	return original.PossibleMetricUnitValues()
}
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return original.PossibleMonitoringStatusValues()
}
func PossibleServiceOwnersAlertNotificationStatusValues() []ServiceOwnersAlertNotificationStatus {
	return original.PossibleServiceOwnersAlertNotificationStatusValues()
}
func PossibleShareStatusValues() []ShareStatus {
	return original.PossibleShareStatusValues()
}
func PossibleSslStatusValues() []SslStatus {
	return original.PossibleSslStatusValues()
}
func PossibleSupportedDeviceCapabilitiesValues() []SupportedDeviceCapabilities {
	return original.PossibleSupportedDeviceCapabilitiesValues()
}
func PossibleTargetTypeValues() []TargetType {
	return original.PossibleTargetTypeValues()
}
func PossibleUpdateOperationValues() []UpdateOperation {
	return original.PossibleUpdateOperationValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
