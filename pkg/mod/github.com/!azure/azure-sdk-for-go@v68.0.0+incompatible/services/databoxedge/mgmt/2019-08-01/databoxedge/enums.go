package databoxedge

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccountType enumerates the values for account type.
type AccountType string

const (
	// BlobStorage ...
	BlobStorage AccountType = "BlobStorage"
	// GeneralPurposeStorage ...
	GeneralPurposeStorage AccountType = "GeneralPurposeStorage"
)

// PossibleAccountTypeValues returns an array of possible values for the AccountType const type.
func PossibleAccountTypeValues() []AccountType {
	return []AccountType{BlobStorage, GeneralPurposeStorage}
}

// AlertSeverity enumerates the values for alert severity.
type AlertSeverity string

const (
	// Critical ...
	Critical AlertSeverity = "Critical"
	// Informational ...
	Informational AlertSeverity = "Informational"
	// Warning ...
	Warning AlertSeverity = "Warning"
)

// PossibleAlertSeverityValues returns an array of possible values for the AlertSeverity const type.
func PossibleAlertSeverityValues() []AlertSeverity {
	return []AlertSeverity{Critical, Informational, Warning}
}

// AuthenticationType enumerates the values for authentication type.
type AuthenticationType string

const (
	// AzureActiveDirectory ...
	AzureActiveDirectory AuthenticationType = "AzureActiveDirectory"
	// Invalid ...
	Invalid AuthenticationType = "Invalid"
)

// PossibleAuthenticationTypeValues returns an array of possible values for the AuthenticationType const type.
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return []AuthenticationType{AzureActiveDirectory, Invalid}
}

// AzureContainerDataFormat enumerates the values for azure container data format.
type AzureContainerDataFormat string

const (
	// AzureFile ...
	AzureFile AzureContainerDataFormat = "AzureFile"
	// BlockBlob ...
	BlockBlob AzureContainerDataFormat = "BlockBlob"
	// PageBlob ...
	PageBlob AzureContainerDataFormat = "PageBlob"
)

// PossibleAzureContainerDataFormatValues returns an array of possible values for the AzureContainerDataFormat const type.
func PossibleAzureContainerDataFormatValues() []AzureContainerDataFormat {
	return []AzureContainerDataFormat{AzureFile, BlockBlob, PageBlob}
}

// ClientPermissionType enumerates the values for client permission type.
type ClientPermissionType string

const (
	// NoAccess ...
	NoAccess ClientPermissionType = "NoAccess"
	// ReadOnly ...
	ReadOnly ClientPermissionType = "ReadOnly"
	// ReadWrite ...
	ReadWrite ClientPermissionType = "ReadWrite"
)

// PossibleClientPermissionTypeValues returns an array of possible values for the ClientPermissionType const type.
func PossibleClientPermissionTypeValues() []ClientPermissionType {
	return []ClientPermissionType{NoAccess, ReadOnly, ReadWrite}
}

// ContainerStatus enumerates the values for container status.
type ContainerStatus string

const (
	// NeedsAttention ...
	NeedsAttention ContainerStatus = "NeedsAttention"
	// Offline ...
	Offline ContainerStatus = "Offline"
	// OK ...
	OK ContainerStatus = "OK"
	// Unknown ...
	Unknown ContainerStatus = "Unknown"
	// Updating ...
	Updating ContainerStatus = "Updating"
)

// PossibleContainerStatusValues returns an array of possible values for the ContainerStatus const type.
func PossibleContainerStatusValues() []ContainerStatus {
	return []ContainerStatus{NeedsAttention, Offline, OK, Unknown, Updating}
}

// DataPolicy enumerates the values for data policy.
type DataPolicy string

const (
	// Cloud ...
	Cloud DataPolicy = "Cloud"
	// Local ...
	Local DataPolicy = "Local"
)

// PossibleDataPolicyValues returns an array of possible values for the DataPolicy const type.
func PossibleDataPolicyValues() []DataPolicy {
	return []DataPolicy{Cloud, Local}
}

// DayOfWeek enumerates the values for day of week.
type DayOfWeek string

const (
	// Friday ...
	Friday DayOfWeek = "Friday"
	// Monday ...
	Monday DayOfWeek = "Monday"
	// Saturday ...
	Saturday DayOfWeek = "Saturday"
	// Sunday ...
	Sunday DayOfWeek = "Sunday"
	// Thursday ...
	Thursday DayOfWeek = "Thursday"
	// Tuesday ...
	Tuesday DayOfWeek = "Tuesday"
	// Wednesday ...
	Wednesday DayOfWeek = "Wednesday"
)

// PossibleDayOfWeekValues returns an array of possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{Friday, Monday, Saturday, Sunday, Thursday, Tuesday, Wednesday}
}

// DeviceStatus enumerates the values for device status.
type DeviceStatus string

const (
	// DeviceStatusDisconnected ...
	DeviceStatusDisconnected DeviceStatus = "Disconnected"
	// DeviceStatusMaintenance ...
	DeviceStatusMaintenance DeviceStatus = "Maintenance"
	// DeviceStatusNeedsAttention ...
	DeviceStatusNeedsAttention DeviceStatus = "NeedsAttention"
	// DeviceStatusOffline ...
	DeviceStatusOffline DeviceStatus = "Offline"
	// DeviceStatusOnline ...
	DeviceStatusOnline DeviceStatus = "Online"
	// DeviceStatusPartiallyDisconnected ...
	DeviceStatusPartiallyDisconnected DeviceStatus = "PartiallyDisconnected"
	// DeviceStatusReadyToSetup ...
	DeviceStatusReadyToSetup DeviceStatus = "ReadyToSetup"
)

// PossibleDeviceStatusValues returns an array of possible values for the DeviceStatus const type.
func PossibleDeviceStatusValues() []DeviceStatus {
	return []DeviceStatus{DeviceStatusDisconnected, DeviceStatusMaintenance, DeviceStatusNeedsAttention, DeviceStatusOffline, DeviceStatusOnline, DeviceStatusPartiallyDisconnected, DeviceStatusReadyToSetup}
}

// DeviceType enumerates the values for device type.
type DeviceType string

const (
	// DataBoxEdgeDevice ...
	DataBoxEdgeDevice DeviceType = "DataBoxEdgeDevice"
)

// PossibleDeviceTypeValues returns an array of possible values for the DeviceType const type.
func PossibleDeviceTypeValues() []DeviceType {
	return []DeviceType{DataBoxEdgeDevice}
}

// DownloadPhase enumerates the values for download phase.
type DownloadPhase string

const (
	// DownloadPhaseDownloading ...
	DownloadPhaseDownloading DownloadPhase = "Downloading"
	// DownloadPhaseInitializing ...
	DownloadPhaseInitializing DownloadPhase = "Initializing"
	// DownloadPhaseUnknown ...
	DownloadPhaseUnknown DownloadPhase = "Unknown"
	// DownloadPhaseVerifying ...
	DownloadPhaseVerifying DownloadPhase = "Verifying"
)

// PossibleDownloadPhaseValues returns an array of possible values for the DownloadPhase const type.
func PossibleDownloadPhaseValues() []DownloadPhase {
	return []DownloadPhase{DownloadPhaseDownloading, DownloadPhaseInitializing, DownloadPhaseUnknown, DownloadPhaseVerifying}
}

// EncryptionAlgorithm enumerates the values for encryption algorithm.
type EncryptionAlgorithm string

const (
	// AES256 ...
	AES256 EncryptionAlgorithm = "AES256"
	// None ...
	None EncryptionAlgorithm = "None"
	// RSAESPKCS1V15 ...
	RSAESPKCS1V15 EncryptionAlgorithm = "RSAES_PKCS1_v_1_5"
)

// PossibleEncryptionAlgorithmValues returns an array of possible values for the EncryptionAlgorithm const type.
func PossibleEncryptionAlgorithmValues() []EncryptionAlgorithm {
	return []EncryptionAlgorithm{AES256, None, RSAESPKCS1V15}
}

// InstallRebootBehavior enumerates the values for install reboot behavior.
type InstallRebootBehavior string

const (
	// NeverReboots ...
	NeverReboots InstallRebootBehavior = "NeverReboots"
	// RequestReboot ...
	RequestReboot InstallRebootBehavior = "RequestReboot"
	// RequiresReboot ...
	RequiresReboot InstallRebootBehavior = "RequiresReboot"
)

// PossibleInstallRebootBehaviorValues returns an array of possible values for the InstallRebootBehavior const type.
func PossibleInstallRebootBehaviorValues() []InstallRebootBehavior {
	return []InstallRebootBehavior{NeverReboots, RequestReboot, RequiresReboot}
}

// JobStatus enumerates the values for job status.
type JobStatus string

const (
	// JobStatusCanceled ...
	JobStatusCanceled JobStatus = "Canceled"
	// JobStatusFailed ...
	JobStatusFailed JobStatus = "Failed"
	// JobStatusInvalid ...
	JobStatusInvalid JobStatus = "Invalid"
	// JobStatusPaused ...
	JobStatusPaused JobStatus = "Paused"
	// JobStatusRunning ...
	JobStatusRunning JobStatus = "Running"
	// JobStatusScheduled ...
	JobStatusScheduled JobStatus = "Scheduled"
	// JobStatusSucceeded ...
	JobStatusSucceeded JobStatus = "Succeeded"
)

// PossibleJobStatusValues returns an array of possible values for the JobStatus const type.
func PossibleJobStatusValues() []JobStatus {
	return []JobStatus{JobStatusCanceled, JobStatusFailed, JobStatusInvalid, JobStatusPaused, JobStatusRunning, JobStatusScheduled, JobStatusSucceeded}
}

// JobType enumerates the values for job type.
type JobType string

const (
	// JobTypeDownloadUpdates ...
	JobTypeDownloadUpdates JobType = "DownloadUpdates"
	// JobTypeInstallUpdates ...
	JobTypeInstallUpdates JobType = "InstallUpdates"
	// JobTypeInvalid ...
	JobTypeInvalid JobType = "Invalid"
	// JobTypeRefreshContainer ...
	JobTypeRefreshContainer JobType = "RefreshContainer"
	// JobTypeRefreshShare ...
	JobTypeRefreshShare JobType = "RefreshShare"
	// JobTypeScanForUpdates ...
	JobTypeScanForUpdates JobType = "ScanForUpdates"
)

// PossibleJobTypeValues returns an array of possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{JobTypeDownloadUpdates, JobTypeInstallUpdates, JobTypeInvalid, JobTypeRefreshContainer, JobTypeRefreshShare, JobTypeScanForUpdates}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindIOT ...
	KindIOT Kind = "IOT"
	// KindRole ...
	KindRole Kind = "Role"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindIOT, KindRole}
}

// KindBasicTrigger enumerates the values for kind basic trigger.
type KindBasicTrigger string

const (
	// KindFileEvent ...
	KindFileEvent KindBasicTrigger = "FileEvent"
	// KindPeriodicTimerEvent ...
	KindPeriodicTimerEvent KindBasicTrigger = "PeriodicTimerEvent"
	// KindTrigger ...
	KindTrigger KindBasicTrigger = "Trigger"
)

// PossibleKindBasicTriggerValues returns an array of possible values for the KindBasicTrigger const type.
func PossibleKindBasicTriggerValues() []KindBasicTrigger {
	return []KindBasicTrigger{KindFileEvent, KindPeriodicTimerEvent, KindTrigger}
}

// MetricAggregationType enumerates the values for metric aggregation type.
type MetricAggregationType string

const (
	// MetricAggregationTypeAverage ...
	MetricAggregationTypeAverage MetricAggregationType = "Average"
	// MetricAggregationTypeCount ...
	MetricAggregationTypeCount MetricAggregationType = "Count"
	// MetricAggregationTypeMaximum ...
	MetricAggregationTypeMaximum MetricAggregationType = "Maximum"
	// MetricAggregationTypeMinimum ...
	MetricAggregationTypeMinimum MetricAggregationType = "Minimum"
	// MetricAggregationTypeNone ...
	MetricAggregationTypeNone MetricAggregationType = "None"
	// MetricAggregationTypeNotSpecified ...
	MetricAggregationTypeNotSpecified MetricAggregationType = "NotSpecified"
	// MetricAggregationTypeTotal ...
	MetricAggregationTypeTotal MetricAggregationType = "Total"
)

// PossibleMetricAggregationTypeValues returns an array of possible values for the MetricAggregationType const type.
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return []MetricAggregationType{MetricAggregationTypeAverage, MetricAggregationTypeCount, MetricAggregationTypeMaximum, MetricAggregationTypeMinimum, MetricAggregationTypeNone, MetricAggregationTypeNotSpecified, MetricAggregationTypeTotal}
}

// MetricCategory enumerates the values for metric category.
type MetricCategory string

const (
	// Capacity ...
	Capacity MetricCategory = "Capacity"
	// Transaction ...
	Transaction MetricCategory = "Transaction"
)

// PossibleMetricCategoryValues returns an array of possible values for the MetricCategory const type.
func PossibleMetricCategoryValues() []MetricCategory {
	return []MetricCategory{Capacity, Transaction}
}

// MetricUnit enumerates the values for metric unit.
type MetricUnit string

const (
	// Bytes ...
	Bytes MetricUnit = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond MetricUnit = "BytesPerSecond"
	// Count ...
	Count MetricUnit = "Count"
	// CountPerSecond ...
	CountPerSecond MetricUnit = "CountPerSecond"
	// Milliseconds ...
	Milliseconds MetricUnit = "Milliseconds"
	// NotSpecified ...
	NotSpecified MetricUnit = "NotSpecified"
	// Percent ...
	Percent MetricUnit = "Percent"
	// Seconds ...
	Seconds MetricUnit = "Seconds"
)

// PossibleMetricUnitValues returns an array of possible values for the MetricUnit const type.
func PossibleMetricUnitValues() []MetricUnit {
	return []MetricUnit{Bytes, BytesPerSecond, Count, CountPerSecond, Milliseconds, NotSpecified, Percent, Seconds}
}

// MonitoringStatus enumerates the values for monitoring status.
type MonitoringStatus string

const (
	// Disabled ...
	Disabled MonitoringStatus = "Disabled"
	// Enabled ...
	Enabled MonitoringStatus = "Enabled"
)

// PossibleMonitoringStatusValues returns an array of possible values for the MonitoringStatus const type.
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return []MonitoringStatus{Disabled, Enabled}
}

// NetworkAdapterDHCPStatus enumerates the values for network adapter dhcp status.
type NetworkAdapterDHCPStatus string

const (
	// NetworkAdapterDHCPStatusDisabled ...
	NetworkAdapterDHCPStatusDisabled NetworkAdapterDHCPStatus = "Disabled"
	// NetworkAdapterDHCPStatusEnabled ...
	NetworkAdapterDHCPStatusEnabled NetworkAdapterDHCPStatus = "Enabled"
)

// PossibleNetworkAdapterDHCPStatusValues returns an array of possible values for the NetworkAdapterDHCPStatus const type.
func PossibleNetworkAdapterDHCPStatusValues() []NetworkAdapterDHCPStatus {
	return []NetworkAdapterDHCPStatus{NetworkAdapterDHCPStatusDisabled, NetworkAdapterDHCPStatusEnabled}
}

// NetworkAdapterRDMAStatus enumerates the values for network adapter rdma status.
type NetworkAdapterRDMAStatus string

const (
	// Capable ...
	Capable NetworkAdapterRDMAStatus = "Capable"
	// Incapable ...
	Incapable NetworkAdapterRDMAStatus = "Incapable"
)

// PossibleNetworkAdapterRDMAStatusValues returns an array of possible values for the NetworkAdapterRDMAStatus const type.
func PossibleNetworkAdapterRDMAStatusValues() []NetworkAdapterRDMAStatus {
	return []NetworkAdapterRDMAStatus{Capable, Incapable}
}

// NetworkAdapterStatus enumerates the values for network adapter status.
type NetworkAdapterStatus string

const (
	// Active ...
	Active NetworkAdapterStatus = "Active"
	// Inactive ...
	Inactive NetworkAdapterStatus = "Inactive"
)

// PossibleNetworkAdapterStatusValues returns an array of possible values for the NetworkAdapterStatus const type.
func PossibleNetworkAdapterStatusValues() []NetworkAdapterStatus {
	return []NetworkAdapterStatus{Active, Inactive}
}

// NetworkGroup enumerates the values for network group.
type NetworkGroup string

const (
	// NetworkGroupNone ...
	NetworkGroupNone NetworkGroup = "None"
	// NetworkGroupNonRDMA ...
	NetworkGroupNonRDMA NetworkGroup = "NonRDMA"
	// NetworkGroupRDMA ...
	NetworkGroupRDMA NetworkGroup = "RDMA"
)

// PossibleNetworkGroupValues returns an array of possible values for the NetworkGroup const type.
func PossibleNetworkGroupValues() []NetworkGroup {
	return []NetworkGroup{NetworkGroupNone, NetworkGroupNonRDMA, NetworkGroupRDMA}
}

// NodeStatus enumerates the values for node status.
type NodeStatus string

const (
	// NodeStatusDown ...
	NodeStatusDown NodeStatus = "Down"
	// NodeStatusRebooting ...
	NodeStatusRebooting NodeStatus = "Rebooting"
	// NodeStatusShuttingDown ...
	NodeStatusShuttingDown NodeStatus = "ShuttingDown"
	// NodeStatusUnknown ...
	NodeStatusUnknown NodeStatus = "Unknown"
	// NodeStatusUp ...
	NodeStatusUp NodeStatus = "Up"
)

// PossibleNodeStatusValues returns an array of possible values for the NodeStatus const type.
func PossibleNodeStatusValues() []NodeStatus {
	return []NodeStatus{NodeStatusDown, NodeStatusRebooting, NodeStatusShuttingDown, NodeStatusUnknown, NodeStatusUp}
}

// OrderState enumerates the values for order state.
type OrderState string

const (
	// Arriving ...
	Arriving OrderState = "Arriving"
	// AwaitingFulfilment ...
	AwaitingFulfilment OrderState = "AwaitingFulfilment"
	// AwaitingPreparation ...
	AwaitingPreparation OrderState = "AwaitingPreparation"
	// AwaitingReturnShipment ...
	AwaitingReturnShipment OrderState = "AwaitingReturnShipment"
	// AwaitingShipment ...
	AwaitingShipment OrderState = "AwaitingShipment"
	// CollectedAtMicrosoft ...
	CollectedAtMicrosoft OrderState = "CollectedAtMicrosoft"
	// Declined ...
	Declined OrderState = "Declined"
	// Delivered ...
	Delivered OrderState = "Delivered"
	// LostDevice ...
	LostDevice OrderState = "LostDevice"
	// ReplacementRequested ...
	ReplacementRequested OrderState = "ReplacementRequested"
	// ReturnInitiated ...
	ReturnInitiated OrderState = "ReturnInitiated"
	// Shipped ...
	Shipped OrderState = "Shipped"
	// ShippedBack ...
	ShippedBack OrderState = "ShippedBack"
	// Untracked ...
	Untracked OrderState = "Untracked"
)

// PossibleOrderStateValues returns an array of possible values for the OrderState const type.
func PossibleOrderStateValues() []OrderState {
	return []OrderState{Arriving, AwaitingFulfilment, AwaitingPreparation, AwaitingReturnShipment, AwaitingShipment, CollectedAtMicrosoft, Declined, Delivered, LostDevice, ReplacementRequested, ReturnInitiated, Shipped, ShippedBack, Untracked}
}

// PlatformType enumerates the values for platform type.
type PlatformType string

const (
	// Linux ...
	Linux PlatformType = "Linux"
	// Windows ...
	Windows PlatformType = "Windows"
)

// PossiblePlatformTypeValues returns an array of possible values for the PlatformType const type.
func PossiblePlatformTypeValues() []PlatformType {
	return []PlatformType{Linux, Windows}
}

// RoleStatus enumerates the values for role status.
type RoleStatus string

const (
	// RoleStatusDisabled ...
	RoleStatusDisabled RoleStatus = "Disabled"
	// RoleStatusEnabled ...
	RoleStatusEnabled RoleStatus = "Enabled"
)

// PossibleRoleStatusValues returns an array of possible values for the RoleStatus const type.
func PossibleRoleStatusValues() []RoleStatus {
	return []RoleStatus{RoleStatusDisabled, RoleStatusEnabled}
}

// RoleTypes enumerates the values for role types.
type RoleTypes string

const (
	// ASA ...
	ASA RoleTypes = "ASA"
	// Cognitive ...
	Cognitive RoleTypes = "Cognitive"
	// Functions ...
	Functions RoleTypes = "Functions"
	// IOT ...
	IOT RoleTypes = "IOT"
)

// PossibleRoleTypesValues returns an array of possible values for the RoleTypes const type.
func PossibleRoleTypesValues() []RoleTypes {
	return []RoleTypes{ASA, Cognitive, Functions, IOT}
}

// ShareAccessProtocol enumerates the values for share access protocol.
type ShareAccessProtocol string

const (
	// NFS ...
	NFS ShareAccessProtocol = "NFS"
	// SMB ...
	SMB ShareAccessProtocol = "SMB"
)

// PossibleShareAccessProtocolValues returns an array of possible values for the ShareAccessProtocol const type.
func PossibleShareAccessProtocolValues() []ShareAccessProtocol {
	return []ShareAccessProtocol{NFS, SMB}
}

// ShareAccessType enumerates the values for share access type.
type ShareAccessType string

const (
	// Change ...
	Change ShareAccessType = "Change"
	// Custom ...
	Custom ShareAccessType = "Custom"
	// Read ...
	Read ShareAccessType = "Read"
)

// PossibleShareAccessTypeValues returns an array of possible values for the ShareAccessType const type.
func PossibleShareAccessTypeValues() []ShareAccessType {
	return []ShareAccessType{Change, Custom, Read}
}

// ShareStatus enumerates the values for share status.
type ShareStatus string

const (
	// ShareStatusNeedsAttention ...
	ShareStatusNeedsAttention ShareStatus = "NeedsAttention"
	// ShareStatusOffline ...
	ShareStatusOffline ShareStatus = "Offline"
	// ShareStatusOK ...
	ShareStatusOK ShareStatus = "OK"
	// ShareStatusUnknown ...
	ShareStatusUnknown ShareStatus = "Unknown"
	// ShareStatusUpdating ...
	ShareStatusUpdating ShareStatus = "Updating"
)

// PossibleShareStatusValues returns an array of possible values for the ShareStatus const type.
func PossibleShareStatusValues() []ShareStatus {
	return []ShareStatus{ShareStatusNeedsAttention, ShareStatusOffline, ShareStatusOK, ShareStatusUnknown, ShareStatusUpdating}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Edge ...
	Edge SkuName = "Edge"
	// Gateway ...
	Gateway SkuName = "Gateway"
	// TEA1Node ...
	TEA1Node SkuName = "TEA_1Node"
	// TEA1NodeHeater ...
	TEA1NodeHeater SkuName = "TEA_1Node_Heater"
	// TEA1NodeUPS ...
	TEA1NodeUPS SkuName = "TEA_1Node_UPS"
	// TEA1NodeUPSHeater ...
	TEA1NodeUPSHeater SkuName = "TEA_1Node_UPS_Heater"
	// TEA4NodeHeater ...
	TEA4NodeHeater SkuName = "TEA_4Node_Heater"
	// TEA4NodeUPSHeater ...
	TEA4NodeUPSHeater SkuName = "TEA_4Node_UPS_Heater"
	// TMA ...
	TMA SkuName = "TMA"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Edge, Gateway, TEA1Node, TEA1NodeHeater, TEA1NodeUPS, TEA1NodeUPSHeater, TEA4NodeHeater, TEA4NodeUPSHeater, TMA}
}

// SkuRestrictionReasonCode enumerates the values for sku restriction reason code.
type SkuRestrictionReasonCode string

const (
	// NotAvailableForSubscription ...
	NotAvailableForSubscription SkuRestrictionReasonCode = "NotAvailableForSubscription"
	// QuotaID ...
	QuotaID SkuRestrictionReasonCode = "QuotaId"
)

// PossibleSkuRestrictionReasonCodeValues returns an array of possible values for the SkuRestrictionReasonCode const type.
func PossibleSkuRestrictionReasonCodeValues() []SkuRestrictionReasonCode {
	return []SkuRestrictionReasonCode{NotAvailableForSubscription, QuotaID}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Standard}
}

// SSLStatus enumerates the values for ssl status.
type SSLStatus string

const (
	// SSLStatusDisabled ...
	SSLStatusDisabled SSLStatus = "Disabled"
	// SSLStatusEnabled ...
	SSLStatusEnabled SSLStatus = "Enabled"
)

// PossibleSSLStatusValues returns an array of possible values for the SSLStatus const type.
func PossibleSSLStatusValues() []SSLStatus {
	return []SSLStatus{SSLStatusDisabled, SSLStatusEnabled}
}

// StorageAccountStatus enumerates the values for storage account status.
type StorageAccountStatus string

const (
	// StorageAccountStatusNeedsAttention ...
	StorageAccountStatusNeedsAttention StorageAccountStatus = "NeedsAttention"
	// StorageAccountStatusOffline ...
	StorageAccountStatusOffline StorageAccountStatus = "Offline"
	// StorageAccountStatusOK ...
	StorageAccountStatusOK StorageAccountStatus = "OK"
	// StorageAccountStatusUnknown ...
	StorageAccountStatusUnknown StorageAccountStatus = "Unknown"
	// StorageAccountStatusUpdating ...
	StorageAccountStatusUpdating StorageAccountStatus = "Updating"
)

// PossibleStorageAccountStatusValues returns an array of possible values for the StorageAccountStatus const type.
func PossibleStorageAccountStatusValues() []StorageAccountStatus {
	return []StorageAccountStatus{StorageAccountStatusNeedsAttention, StorageAccountStatusOffline, StorageAccountStatusOK, StorageAccountStatusUnknown, StorageAccountStatusUpdating}
}

// TimeGrain enumerates the values for time grain.
type TimeGrain string

const (
	// PT12H ...
	PT12H TimeGrain = "PT12H"
	// PT15M ...
	PT15M TimeGrain = "PT15M"
	// PT1D ...
	PT1D TimeGrain = "PT1D"
	// PT1H ...
	PT1H TimeGrain = "PT1H"
	// PT1M ...
	PT1M TimeGrain = "PT1M"
	// PT30M ...
	PT30M TimeGrain = "PT30M"
	// PT5M ...
	PT5M TimeGrain = "PT5M"
	// PT6H ...
	PT6H TimeGrain = "PT6H"
)

// PossibleTimeGrainValues returns an array of possible values for the TimeGrain const type.
func PossibleTimeGrainValues() []TimeGrain {
	return []TimeGrain{PT12H, PT15M, PT1D, PT1H, PT1M, PT30M, PT5M, PT6H}
}

// UpdateOperation enumerates the values for update operation.
type UpdateOperation string

const (
	// UpdateOperationDownload ...
	UpdateOperationDownload UpdateOperation = "Download"
	// UpdateOperationInstall ...
	UpdateOperationInstall UpdateOperation = "Install"
	// UpdateOperationNone ...
	UpdateOperationNone UpdateOperation = "None"
	// UpdateOperationScan ...
	UpdateOperationScan UpdateOperation = "Scan"
)

// PossibleUpdateOperationValues returns an array of possible values for the UpdateOperation const type.
func PossibleUpdateOperationValues() []UpdateOperation {
	return []UpdateOperation{UpdateOperationDownload, UpdateOperationInstall, UpdateOperationNone, UpdateOperationScan}
}

// UpdateOperationStage enumerates the values for update operation stage.
type UpdateOperationStage string

const (
	// UpdateOperationStageDownloadComplete ...
	UpdateOperationStageDownloadComplete UpdateOperationStage = "DownloadComplete"
	// UpdateOperationStageDownloadFailed ...
	UpdateOperationStageDownloadFailed UpdateOperationStage = "DownloadFailed"
	// UpdateOperationStageDownloadStarted ...
	UpdateOperationStageDownloadStarted UpdateOperationStage = "DownloadStarted"
	// UpdateOperationStageFailure ...
	UpdateOperationStageFailure UpdateOperationStage = "Failure"
	// UpdateOperationStageInitial ...
	UpdateOperationStageInitial UpdateOperationStage = "Initial"
	// UpdateOperationStageInstallComplete ...
	UpdateOperationStageInstallComplete UpdateOperationStage = "InstallComplete"
	// UpdateOperationStageInstallFailed ...
	UpdateOperationStageInstallFailed UpdateOperationStage = "InstallFailed"
	// UpdateOperationStageInstallStarted ...
	UpdateOperationStageInstallStarted UpdateOperationStage = "InstallStarted"
	// UpdateOperationStageRebootInitiated ...
	UpdateOperationStageRebootInitiated UpdateOperationStage = "RebootInitiated"
	// UpdateOperationStageRescanComplete ...
	UpdateOperationStageRescanComplete UpdateOperationStage = "RescanComplete"
	// UpdateOperationStageRescanFailed ...
	UpdateOperationStageRescanFailed UpdateOperationStage = "RescanFailed"
	// UpdateOperationStageRescanStarted ...
	UpdateOperationStageRescanStarted UpdateOperationStage = "RescanStarted"
	// UpdateOperationStageScanComplete ...
	UpdateOperationStageScanComplete UpdateOperationStage = "ScanComplete"
	// UpdateOperationStageScanFailed ...
	UpdateOperationStageScanFailed UpdateOperationStage = "ScanFailed"
	// UpdateOperationStageScanStarted ...
	UpdateOperationStageScanStarted UpdateOperationStage = "ScanStarted"
	// UpdateOperationStageSuccess ...
	UpdateOperationStageSuccess UpdateOperationStage = "Success"
	// UpdateOperationStageUnknown ...
	UpdateOperationStageUnknown UpdateOperationStage = "Unknown"
)

// PossibleUpdateOperationStageValues returns an array of possible values for the UpdateOperationStage const type.
func PossibleUpdateOperationStageValues() []UpdateOperationStage {
	return []UpdateOperationStage{UpdateOperationStageDownloadComplete, UpdateOperationStageDownloadFailed, UpdateOperationStageDownloadStarted, UpdateOperationStageFailure, UpdateOperationStageInitial, UpdateOperationStageInstallComplete, UpdateOperationStageInstallFailed, UpdateOperationStageInstallStarted, UpdateOperationStageRebootInitiated, UpdateOperationStageRescanComplete, UpdateOperationStageRescanFailed, UpdateOperationStageRescanStarted, UpdateOperationStageScanComplete, UpdateOperationStageScanFailed, UpdateOperationStageScanStarted, UpdateOperationStageSuccess, UpdateOperationStageUnknown}
}

// UserType enumerates the values for user type.
type UserType string

const (
	// UserTypeARM ...
	UserTypeARM UserType = "ARM"
	// UserTypeLocalManagement ...
	UserTypeLocalManagement UserType = "LocalManagement"
	// UserTypeShare ...
	UserTypeShare UserType = "Share"
)

// PossibleUserTypeValues returns an array of possible values for the UserType const type.
func PossibleUserTypeValues() []UserType {
	return []UserType{UserTypeARM, UserTypeLocalManagement, UserTypeShare}
}