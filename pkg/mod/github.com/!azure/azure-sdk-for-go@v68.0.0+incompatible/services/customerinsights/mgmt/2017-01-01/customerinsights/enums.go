package customerinsights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CalculationWindowTypes enumerates the values for calculation window types.
type CalculationWindowTypes string

const (
	// Day ...
	Day CalculationWindowTypes = "Day"
	// Hour ...
	Hour CalculationWindowTypes = "Hour"
	// Lifetime ...
	Lifetime CalculationWindowTypes = "Lifetime"
	// Month ...
	Month CalculationWindowTypes = "Month"
	// Week ...
	Week CalculationWindowTypes = "Week"
)

// PossibleCalculationWindowTypesValues returns an array of possible values for the CalculationWindowTypes const type.
func PossibleCalculationWindowTypesValues() []CalculationWindowTypes {
	return []CalculationWindowTypes{Day, Hour, Lifetime, Month, Week}
}

// CardinalityTypes enumerates the values for cardinality types.
type CardinalityTypes string

const (
	// ManyToMany ...
	ManyToMany CardinalityTypes = "ManyToMany"
	// OneToMany ...
	OneToMany CardinalityTypes = "OneToMany"
	// OneToOne ...
	OneToOne CardinalityTypes = "OneToOne"
)

// PossibleCardinalityTypesValues returns an array of possible values for the CardinalityTypes const type.
func PossibleCardinalityTypesValues() []CardinalityTypes {
	return []CardinalityTypes{ManyToMany, OneToMany, OneToOne}
}

// CompletionOperationTypes enumerates the values for completion operation types.
type CompletionOperationTypes string

const (
	// DeleteFile ...
	DeleteFile CompletionOperationTypes = "DeleteFile"
	// DoNothing ...
	DoNothing CompletionOperationTypes = "DoNothing"
	// MoveFile ...
	MoveFile CompletionOperationTypes = "MoveFile"
)

// PossibleCompletionOperationTypesValues returns an array of possible values for the CompletionOperationTypes const type.
func PossibleCompletionOperationTypesValues() []CompletionOperationTypes {
	return []CompletionOperationTypes{DeleteFile, DoNothing, MoveFile}
}

// ConnectorMappingStates enumerates the values for connector mapping states.
type ConnectorMappingStates string

const (
	// Created ...
	Created ConnectorMappingStates = "Created"
	// Creating ...
	Creating ConnectorMappingStates = "Creating"
	// Expiring ...
	Expiring ConnectorMappingStates = "Expiring"
	// Failed ...
	Failed ConnectorMappingStates = "Failed"
	// Ready ...
	Ready ConnectorMappingStates = "Ready"
	// Running ...
	Running ConnectorMappingStates = "Running"
	// Stopped ...
	Stopped ConnectorMappingStates = "Stopped"
)

// PossibleConnectorMappingStatesValues returns an array of possible values for the ConnectorMappingStates const type.
func PossibleConnectorMappingStatesValues() []ConnectorMappingStates {
	return []ConnectorMappingStates{Created, Creating, Expiring, Failed, Ready, Running, Stopped}
}

// ConnectorStates enumerates the values for connector states.
type ConnectorStates string

const (
	// ConnectorStatesCreated ...
	ConnectorStatesCreated ConnectorStates = "Created"
	// ConnectorStatesCreating ...
	ConnectorStatesCreating ConnectorStates = "Creating"
	// ConnectorStatesDeleting ...
	ConnectorStatesDeleting ConnectorStates = "Deleting"
	// ConnectorStatesExpiring ...
	ConnectorStatesExpiring ConnectorStates = "Expiring"
	// ConnectorStatesFailed ...
	ConnectorStatesFailed ConnectorStates = "Failed"
	// ConnectorStatesReady ...
	ConnectorStatesReady ConnectorStates = "Ready"
)

// PossibleConnectorStatesValues returns an array of possible values for the ConnectorStates const type.
func PossibleConnectorStatesValues() []ConnectorStates {
	return []ConnectorStates{ConnectorStatesCreated, ConnectorStatesCreating, ConnectorStatesDeleting, ConnectorStatesExpiring, ConnectorStatesFailed, ConnectorStatesReady}
}

// ConnectorTypes enumerates the values for connector types.
type ConnectorTypes string

const (
	// AzureBlob ...
	AzureBlob ConnectorTypes = "AzureBlob"
	// CRM ...
	CRM ConnectorTypes = "CRM"
	// ExchangeOnline ...
	ExchangeOnline ConnectorTypes = "ExchangeOnline"
	// None ...
	None ConnectorTypes = "None"
	// Outbound ...
	Outbound ConnectorTypes = "Outbound"
	// Salesforce ...
	Salesforce ConnectorTypes = "Salesforce"
)

// PossibleConnectorTypesValues returns an array of possible values for the ConnectorTypes const type.
func PossibleConnectorTypesValues() []ConnectorTypes {
	return []ConnectorTypes{AzureBlob, CRM, ExchangeOnline, None, Outbound, Salesforce}
}

// DataSourceType enumerates the values for data source type.
type DataSourceType string

const (
	// DataSourceTypeConnector ...
	DataSourceTypeConnector DataSourceType = "Connector"
	// DataSourceTypeLinkInteraction ...
	DataSourceTypeLinkInteraction DataSourceType = "LinkInteraction"
	// DataSourceTypeSystemDefault ...
	DataSourceTypeSystemDefault DataSourceType = "SystemDefault"
)

// PossibleDataSourceTypeValues returns an array of possible values for the DataSourceType const type.
func PossibleDataSourceTypeValues() []DataSourceType {
	return []DataSourceType{DataSourceTypeConnector, DataSourceTypeLinkInteraction, DataSourceTypeSystemDefault}
}

// EntityTypes enumerates the values for entity types.
type EntityTypes string

const (
	// EntityTypesInteraction ...
	EntityTypesInteraction EntityTypes = "Interaction"
	// EntityTypesNone ...
	EntityTypesNone EntityTypes = "None"
	// EntityTypesProfile ...
	EntityTypesProfile EntityTypes = "Profile"
	// EntityTypesRelationship ...
	EntityTypesRelationship EntityTypes = "Relationship"
)

// PossibleEntityTypesValues returns an array of possible values for the EntityTypes const type.
func PossibleEntityTypesValues() []EntityTypes {
	return []EntityTypes{EntityTypesInteraction, EntityTypesNone, EntityTypesProfile, EntityTypesRelationship}
}

// ErrorManagementTypes enumerates the values for error management types.
type ErrorManagementTypes string

const (
	// RejectAndContinue ...
	RejectAndContinue ErrorManagementTypes = "RejectAndContinue"
	// RejectUntilLimit ...
	RejectUntilLimit ErrorManagementTypes = "RejectUntilLimit"
	// StopImport ...
	StopImport ErrorManagementTypes = "StopImport"
)

// PossibleErrorManagementTypesValues returns an array of possible values for the ErrorManagementTypes const type.
func PossibleErrorManagementTypesValues() []ErrorManagementTypes {
	return []ErrorManagementTypes{RejectAndContinue, RejectUntilLimit, StopImport}
}

// FrequencyTypes enumerates the values for frequency types.
type FrequencyTypes string

const (
	// FrequencyTypesDay ...
	FrequencyTypesDay FrequencyTypes = "Day"
	// FrequencyTypesHour ...
	FrequencyTypesHour FrequencyTypes = "Hour"
	// FrequencyTypesMinute ...
	FrequencyTypesMinute FrequencyTypes = "Minute"
	// FrequencyTypesMonth ...
	FrequencyTypesMonth FrequencyTypes = "Month"
	// FrequencyTypesWeek ...
	FrequencyTypesWeek FrequencyTypes = "Week"
)

// PossibleFrequencyTypesValues returns an array of possible values for the FrequencyTypes const type.
func PossibleFrequencyTypesValues() []FrequencyTypes {
	return []FrequencyTypes{FrequencyTypesDay, FrequencyTypesHour, FrequencyTypesMinute, FrequencyTypesMonth, FrequencyTypesWeek}
}

// InstanceOperationType enumerates the values for instance operation type.
type InstanceOperationType string

const (
	// Delete ...
	Delete InstanceOperationType = "Delete"
	// Upsert ...
	Upsert InstanceOperationType = "Upsert"
)

// PossibleInstanceOperationTypeValues returns an array of possible values for the InstanceOperationType const type.
func PossibleInstanceOperationTypeValues() []InstanceOperationType {
	return []InstanceOperationType{Delete, Upsert}
}

// KpiFunctions enumerates the values for kpi functions.
type KpiFunctions string

const (
	// KpiFunctionsAvg ...
	KpiFunctionsAvg KpiFunctions = "Avg"
	// KpiFunctionsCount ...
	KpiFunctionsCount KpiFunctions = "Count"
	// KpiFunctionsCountDistinct ...
	KpiFunctionsCountDistinct KpiFunctions = "CountDistinct"
	// KpiFunctionsLast ...
	KpiFunctionsLast KpiFunctions = "Last"
	// KpiFunctionsMax ...
	KpiFunctionsMax KpiFunctions = "Max"
	// KpiFunctionsMin ...
	KpiFunctionsMin KpiFunctions = "Min"
	// KpiFunctionsNone ...
	KpiFunctionsNone KpiFunctions = "None"
	// KpiFunctionsSum ...
	KpiFunctionsSum KpiFunctions = "Sum"
)

// PossibleKpiFunctionsValues returns an array of possible values for the KpiFunctions const type.
func PossibleKpiFunctionsValues() []KpiFunctions {
	return []KpiFunctions{KpiFunctionsAvg, KpiFunctionsCount, KpiFunctionsCountDistinct, KpiFunctionsLast, KpiFunctionsMax, KpiFunctionsMin, KpiFunctionsNone, KpiFunctionsSum}
}

// LinkTypes enumerates the values for link types.
type LinkTypes string

const (
	// CopyIfNull ...
	CopyIfNull LinkTypes = "CopyIfNull"
	// UpdateAlways ...
	UpdateAlways LinkTypes = "UpdateAlways"
)

// PossibleLinkTypesValues returns an array of possible values for the LinkTypes const type.
func PossibleLinkTypesValues() []LinkTypes {
	return []LinkTypes{CopyIfNull, UpdateAlways}
}

// PermissionTypes enumerates the values for permission types.
type PermissionTypes string

const (
	// Manage ...
	Manage PermissionTypes = "Manage"
	// Read ...
	Read PermissionTypes = "Read"
	// Write ...
	Write PermissionTypes = "Write"
)

// PossiblePermissionTypesValues returns an array of possible values for the PermissionTypes const type.
func PossiblePermissionTypesValues() []PermissionTypes {
	return []PermissionTypes{Manage, Read, Write}
}

// ProvisioningStates enumerates the values for provisioning states.
type ProvisioningStates string

const (
	// ProvisioningStatesDeleting ...
	ProvisioningStatesDeleting ProvisioningStates = "Deleting"
	// ProvisioningStatesExpiring ...
	ProvisioningStatesExpiring ProvisioningStates = "Expiring"
	// ProvisioningStatesFailed ...
	ProvisioningStatesFailed ProvisioningStates = "Failed"
	// ProvisioningStatesHumanIntervention ...
	ProvisioningStatesHumanIntervention ProvisioningStates = "HumanIntervention"
	// ProvisioningStatesProvisioning ...
	ProvisioningStatesProvisioning ProvisioningStates = "Provisioning"
	// ProvisioningStatesSucceeded ...
	ProvisioningStatesSucceeded ProvisioningStates = "Succeeded"
)

// PossibleProvisioningStatesValues returns an array of possible values for the ProvisioningStates const type.
func PossibleProvisioningStatesValues() []ProvisioningStates {
	return []ProvisioningStates{ProvisioningStatesDeleting, ProvisioningStatesExpiring, ProvisioningStatesFailed, ProvisioningStatesHumanIntervention, ProvisioningStatesProvisioning, ProvisioningStatesSucceeded}
}

// RoleTypes enumerates the values for role types.
type RoleTypes string

const (
	// Admin ...
	Admin RoleTypes = "Admin"
	// DataAdmin ...
	DataAdmin RoleTypes = "DataAdmin"
	// DataReader ...
	DataReader RoleTypes = "DataReader"
	// ManageAdmin ...
	ManageAdmin RoleTypes = "ManageAdmin"
	// ManageReader ...
	ManageReader RoleTypes = "ManageReader"
	// Reader ...
	Reader RoleTypes = "Reader"
)

// PossibleRoleTypesValues returns an array of possible values for the RoleTypes const type.
func PossibleRoleTypesValues() []RoleTypes {
	return []RoleTypes{Admin, DataAdmin, DataReader, ManageAdmin, ManageReader, Reader}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusActive ...
	StatusActive Status = "Active"
	// StatusDeleted ...
	StatusDeleted Status = "Deleted"
	// StatusNone ...
	StatusNone Status = "None"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusActive, StatusDeleted, StatusNone}
}
