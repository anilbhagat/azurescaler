//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package maintenance

import original "github.com/Azure/azure-sdk-for-go/services/preview/maintenance/mgmt/2022-07-01-preview/maintenance"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type ImpactType = original.ImpactType

const (
	Freeze   ImpactType = original.Freeze
	None     ImpactType = original.None
	Redeploy ImpactType = original.Redeploy
	Restart  ImpactType = original.Restart
)

type RebootOptions = original.RebootOptions

const (
	Always     RebootOptions = original.Always
	IfRequired RebootOptions = original.IfRequired
	Never      RebootOptions = original.Never
)

type Scope = original.Scope

const (
	ScopeExtension          Scope = original.ScopeExtension
	ScopeHost               Scope = original.ScopeHost
	ScopeInGuestPatch       Scope = original.ScopeInGuestPatch
	ScopeOSImage            Scope = original.ScopeOSImage
	ScopeResource           Scope = original.ScopeResource
	ScopeSQLDB              Scope = original.ScopeSQLDB
	ScopeSQLManagedInstance Scope = original.ScopeSQLManagedInstance
)

type TaskScope = original.TaskScope

const (
	TaskScopeGlobal   TaskScope = original.TaskScopeGlobal
	TaskScopeResource TaskScope = original.TaskScopeResource
)

type UpdateStatus = original.UpdateStatus

const (
	Completed  UpdateStatus = original.Completed
	InProgress UpdateStatus = original.InProgress
	Pending    UpdateStatus = original.Pending
	RetryLater UpdateStatus = original.RetryLater
	RetryNow   UpdateStatus = original.RetryNow
)

type Visibility = original.Visibility

const (
	Custom Visibility = original.Custom
	Public Visibility = original.Public
)

type ApplyUpdate = original.ApplyUpdate
type ApplyUpdateForResourceGroupClient = original.ApplyUpdateForResourceGroupClient
type ApplyUpdateProperties = original.ApplyUpdateProperties
type ApplyUpdatesClient = original.ApplyUpdatesClient
type BaseClient = original.BaseClient
type Configuration = original.Configuration
type ConfigurationAssignment = original.ConfigurationAssignment
type ConfigurationAssignmentProperties = original.ConfigurationAssignmentProperties
type ConfigurationAssignmentsClient = original.ConfigurationAssignmentsClient
type ConfigurationAssignmentsWithinSubscriptionClient = original.ConfigurationAssignmentsWithinSubscriptionClient
type ConfigurationProperties = original.ConfigurationProperties
type ConfigurationsClient = original.ConfigurationsClient
type ConfigurationsForResourceGroupClient = original.ConfigurationsForResourceGroupClient
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type InputLinuxParameters = original.InputLinuxParameters
type InputPatchConfiguration = original.InputPatchConfiguration
type InputWindowsParameters = original.InputWindowsParameters
type ListApplyUpdate = original.ListApplyUpdate
type ListConfigurationAssignmentsResult = original.ListConfigurationAssignmentsResult
type ListMaintenanceConfigurationsResult = original.ListMaintenanceConfigurationsResult
type ListUpdatesResult = original.ListUpdatesResult
type Operation = original.Operation
type OperationInfo = original.OperationInfo
type OperationsClient = original.OperationsClient
type OperationsListResult = original.OperationsListResult
type PublicMaintenanceConfigurationsClient = original.PublicMaintenanceConfigurationsClient
type Resource = original.Resource
type SoftwareUpdateConfigurationTasks = original.SoftwareUpdateConfigurationTasks
type SystemData = original.SystemData
type TaskProperties = original.TaskProperties
type Update = original.Update
type UpdateProperties = original.UpdateProperties
type UpdatesClient = original.UpdatesClient
type Window = original.Window

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplyUpdateForResourceGroupClient(subscriptionID string) ApplyUpdateForResourceGroupClient {
	return original.NewApplyUpdateForResourceGroupClient(subscriptionID)
}
func NewApplyUpdateForResourceGroupClientWithBaseURI(baseURI string, subscriptionID string) ApplyUpdateForResourceGroupClient {
	return original.NewApplyUpdateForResourceGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplyUpdatesClient(subscriptionID string) ApplyUpdatesClient {
	return original.NewApplyUpdatesClient(subscriptionID)
}
func NewApplyUpdatesClientWithBaseURI(baseURI string, subscriptionID string) ApplyUpdatesClient {
	return original.NewApplyUpdatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationAssignmentsClient(subscriptionID string) ConfigurationAssignmentsClient {
	return original.NewConfigurationAssignmentsClient(subscriptionID)
}
func NewConfigurationAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationAssignmentsClient {
	return original.NewConfigurationAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationAssignmentsWithinSubscriptionClient(subscriptionID string) ConfigurationAssignmentsWithinSubscriptionClient {
	return original.NewConfigurationAssignmentsWithinSubscriptionClient(subscriptionID)
}
func NewConfigurationAssignmentsWithinSubscriptionClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationAssignmentsWithinSubscriptionClient {
	return original.NewConfigurationAssignmentsWithinSubscriptionClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsForResourceGroupClient(subscriptionID string) ConfigurationsForResourceGroupClient {
	return original.NewConfigurationsForResourceGroupClient(subscriptionID)
}
func NewConfigurationsForResourceGroupClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsForResourceGroupClient {
	return original.NewConfigurationsForResourceGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPublicMaintenanceConfigurationsClient(subscriptionID string) PublicMaintenanceConfigurationsClient {
	return original.NewPublicMaintenanceConfigurationsClient(subscriptionID)
}
func NewPublicMaintenanceConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) PublicMaintenanceConfigurationsClient {
	return original.NewPublicMaintenanceConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUpdatesClient(subscriptionID string) UpdatesClient {
	return original.NewUpdatesClient(subscriptionID)
}
func NewUpdatesClientWithBaseURI(baseURI string, subscriptionID string) UpdatesClient {
	return original.NewUpdatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleImpactTypeValues() []ImpactType {
	return original.PossibleImpactTypeValues()
}
func PossibleRebootOptionsValues() []RebootOptions {
	return original.PossibleRebootOptionsValues()
}
func PossibleScopeValues() []Scope {
	return original.PossibleScopeValues()
}
func PossibleTaskScopeValues() []TaskScope {
	return original.PossibleTaskScopeValues()
}
func PossibleUpdateStatusValues() []UpdateStatus {
	return original.PossibleUpdateStatusValues()
}
func PossibleVisibilityValues() []Visibility {
	return original.PossibleVisibilityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
