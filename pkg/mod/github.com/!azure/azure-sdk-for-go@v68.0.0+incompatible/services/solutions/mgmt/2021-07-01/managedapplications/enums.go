package managedapplications

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionType enumerates the values for action type.
type ActionType string

const (
	// Internal ...
	Internal ActionType = "Internal"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{Internal}
}

// ApplicationArtifactName enumerates the values for application artifact name.
type ApplicationArtifactName string

const (
	// Authorizations ...
	Authorizations ApplicationArtifactName = "Authorizations"
	// CustomRoleDefinition ...
	CustomRoleDefinition ApplicationArtifactName = "CustomRoleDefinition"
	// NotSpecified ...
	NotSpecified ApplicationArtifactName = "NotSpecified"
	// ViewDefinition ...
	ViewDefinition ApplicationArtifactName = "ViewDefinition"
)

// PossibleApplicationArtifactNameValues returns an array of possible values for the ApplicationArtifactName const type.
func PossibleApplicationArtifactNameValues() []ApplicationArtifactName {
	return []ApplicationArtifactName{Authorizations, CustomRoleDefinition, NotSpecified, ViewDefinition}
}

// ApplicationArtifactType enumerates the values for application artifact type.
type ApplicationArtifactType string

const (
	// ApplicationArtifactTypeCustom ...
	ApplicationArtifactTypeCustom ApplicationArtifactType = "Custom"
	// ApplicationArtifactTypeNotSpecified ...
	ApplicationArtifactTypeNotSpecified ApplicationArtifactType = "NotSpecified"
	// ApplicationArtifactTypeTemplate ...
	ApplicationArtifactTypeTemplate ApplicationArtifactType = "Template"
)

// PossibleApplicationArtifactTypeValues returns an array of possible values for the ApplicationArtifactType const type.
func PossibleApplicationArtifactTypeValues() []ApplicationArtifactType {
	return []ApplicationArtifactType{ApplicationArtifactTypeCustom, ApplicationArtifactTypeNotSpecified, ApplicationArtifactTypeTemplate}
}

// ApplicationDefinitionArtifactName enumerates the values for application definition artifact name.
type ApplicationDefinitionArtifactName string

const (
	// ApplicationDefinitionArtifactNameApplicationResourceTemplate ...
	ApplicationDefinitionArtifactNameApplicationResourceTemplate ApplicationDefinitionArtifactName = "ApplicationResourceTemplate"
	// ApplicationDefinitionArtifactNameCreateUIDefinition ...
	ApplicationDefinitionArtifactNameCreateUIDefinition ApplicationDefinitionArtifactName = "CreateUiDefinition"
	// ApplicationDefinitionArtifactNameMainTemplateParameters ...
	ApplicationDefinitionArtifactNameMainTemplateParameters ApplicationDefinitionArtifactName = "MainTemplateParameters"
	// ApplicationDefinitionArtifactNameNotSpecified ...
	ApplicationDefinitionArtifactNameNotSpecified ApplicationDefinitionArtifactName = "NotSpecified"
)

// PossibleApplicationDefinitionArtifactNameValues returns an array of possible values for the ApplicationDefinitionArtifactName const type.
func PossibleApplicationDefinitionArtifactNameValues() []ApplicationDefinitionArtifactName {
	return []ApplicationDefinitionArtifactName{ApplicationDefinitionArtifactNameApplicationResourceTemplate, ApplicationDefinitionArtifactNameCreateUIDefinition, ApplicationDefinitionArtifactNameMainTemplateParameters, ApplicationDefinitionArtifactNameNotSpecified}
}

// ApplicationLockLevel enumerates the values for application lock level.
type ApplicationLockLevel string

const (
	// CanNotDelete ...
	CanNotDelete ApplicationLockLevel = "CanNotDelete"
	// None ...
	None ApplicationLockLevel = "None"
	// ReadOnly ...
	ReadOnly ApplicationLockLevel = "ReadOnly"
)

// PossibleApplicationLockLevelValues returns an array of possible values for the ApplicationLockLevel const type.
func PossibleApplicationLockLevelValues() []ApplicationLockLevel {
	return []ApplicationLockLevel{CanNotDelete, None, ReadOnly}
}

// ApplicationManagementMode enumerates the values for application management mode.
type ApplicationManagementMode string

const (
	// ApplicationManagementModeManaged ...
	ApplicationManagementModeManaged ApplicationManagementMode = "Managed"
	// ApplicationManagementModeNotSpecified ...
	ApplicationManagementModeNotSpecified ApplicationManagementMode = "NotSpecified"
	// ApplicationManagementModeUnmanaged ...
	ApplicationManagementModeUnmanaged ApplicationManagementMode = "Unmanaged"
)

// PossibleApplicationManagementModeValues returns an array of possible values for the ApplicationManagementMode const type.
func PossibleApplicationManagementModeValues() []ApplicationManagementMode {
	return []ApplicationManagementMode{ApplicationManagementModeManaged, ApplicationManagementModeNotSpecified, ApplicationManagementModeUnmanaged}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// DeploymentMode enumerates the values for deployment mode.
type DeploymentMode string

const (
	// DeploymentModeComplete ...
	DeploymentModeComplete DeploymentMode = "Complete"
	// DeploymentModeIncremental ...
	DeploymentModeIncremental DeploymentMode = "Incremental"
	// DeploymentModeNotSpecified ...
	DeploymentModeNotSpecified DeploymentMode = "NotSpecified"
)

// PossibleDeploymentModeValues returns an array of possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{DeploymentModeComplete, DeploymentModeIncremental, DeploymentModeNotSpecified}
}

// JitApprovalMode enumerates the values for jit approval mode.
type JitApprovalMode string

const (
	// JitApprovalModeAutoApprove ...
	JitApprovalModeAutoApprove JitApprovalMode = "AutoApprove"
	// JitApprovalModeManualApprove ...
	JitApprovalModeManualApprove JitApprovalMode = "ManualApprove"
	// JitApprovalModeNotSpecified ...
	JitApprovalModeNotSpecified JitApprovalMode = "NotSpecified"
)

// PossibleJitApprovalModeValues returns an array of possible values for the JitApprovalMode const type.
func PossibleJitApprovalModeValues() []JitApprovalMode {
	return []JitApprovalMode{JitApprovalModeAutoApprove, JitApprovalModeManualApprove, JitApprovalModeNotSpecified}
}

// JitApproverType enumerates the values for jit approver type.
type JitApproverType string

const (
	// Group ...
	Group JitApproverType = "group"
	// User ...
	User JitApproverType = "user"
)

// PossibleJitApproverTypeValues returns an array of possible values for the JitApproverType const type.
func PossibleJitApproverTypeValues() []JitApproverType {
	return []JitApproverType{Group, User}
}

// JitRequestState enumerates the values for jit request state.
type JitRequestState string

const (
	// JitRequestStateApproved ...
	JitRequestStateApproved JitRequestState = "Approved"
	// JitRequestStateCanceled ...
	JitRequestStateCanceled JitRequestState = "Canceled"
	// JitRequestStateDenied ...
	JitRequestStateDenied JitRequestState = "Denied"
	// JitRequestStateExpired ...
	JitRequestStateExpired JitRequestState = "Expired"
	// JitRequestStateFailed ...
	JitRequestStateFailed JitRequestState = "Failed"
	// JitRequestStateNotSpecified ...
	JitRequestStateNotSpecified JitRequestState = "NotSpecified"
	// JitRequestStatePending ...
	JitRequestStatePending JitRequestState = "Pending"
	// JitRequestStateTimeout ...
	JitRequestStateTimeout JitRequestState = "Timeout"
)

// PossibleJitRequestStateValues returns an array of possible values for the JitRequestState const type.
func PossibleJitRequestStateValues() []JitRequestState {
	return []JitRequestState{JitRequestStateApproved, JitRequestStateCanceled, JitRequestStateDenied, JitRequestStateExpired, JitRequestStateFailed, JitRequestStateNotSpecified, JitRequestStatePending, JitRequestStateTimeout}
}

// JitSchedulingType enumerates the values for jit scheduling type.
type JitSchedulingType string

const (
	// JitSchedulingTypeNotSpecified ...
	JitSchedulingTypeNotSpecified JitSchedulingType = "NotSpecified"
	// JitSchedulingTypeOnce ...
	JitSchedulingTypeOnce JitSchedulingType = "Once"
	// JitSchedulingTypeRecurring ...
	JitSchedulingTypeRecurring JitSchedulingType = "Recurring"
)

// PossibleJitSchedulingTypeValues returns an array of possible values for the JitSchedulingType const type.
func PossibleJitSchedulingTypeValues() []JitSchedulingType {
	return []JitSchedulingType{JitSchedulingTypeNotSpecified, JitSchedulingTypeOnce, JitSchedulingTypeRecurring}
}

// Origin enumerates the values for origin.
type Origin string

const (
	// OriginSystem ...
	OriginSystem Origin = "system"
	// OriginUser ...
	OriginUser Origin = "user"
	// OriginUsersystem ...
	OriginUsersystem Origin = "user,system"
)

// PossibleOriginValues returns an array of possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{OriginSystem, OriginUser, OriginUsersystem}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateAccepted ...
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateNotSpecified ...
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	// ProvisioningStateRunning ...
	ProvisioningStateRunning ProvisioningState = "Running"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateAccepted, ProvisioningStateCanceled, ProvisioningStateDeleted, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateNotSpecified, ProvisioningStateRunning, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned ...
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeSystemAssignedUserAssigned ...
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeSystemAssignedUserAssigned, ResourceIdentityTypeUserAssigned}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusElevate ...
	StatusElevate Status = "Elevate"
	// StatusNotSpecified ...
	StatusNotSpecified Status = "NotSpecified"
	// StatusRemove ...
	StatusRemove Status = "Remove"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusElevate, StatusNotSpecified, StatusRemove}
}

// Substatus enumerates the values for substatus.
type Substatus string

const (
	// SubstatusApproved ...
	SubstatusApproved Substatus = "Approved"
	// SubstatusDenied ...
	SubstatusDenied Substatus = "Denied"
	// SubstatusExpired ...
	SubstatusExpired Substatus = "Expired"
	// SubstatusFailed ...
	SubstatusFailed Substatus = "Failed"
	// SubstatusNotSpecified ...
	SubstatusNotSpecified Substatus = "NotSpecified"
	// SubstatusTimeout ...
	SubstatusTimeout Substatus = "Timeout"
)

// PossibleSubstatusValues returns an array of possible values for the Substatus const type.
func PossibleSubstatusValues() []Substatus {
	return []Substatus{SubstatusApproved, SubstatusDenied, SubstatusExpired, SubstatusFailed, SubstatusNotSpecified, SubstatusTimeout}
}
