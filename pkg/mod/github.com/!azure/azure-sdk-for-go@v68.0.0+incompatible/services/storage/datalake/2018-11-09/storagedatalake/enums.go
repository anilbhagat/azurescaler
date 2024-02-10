package storagedatalake

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// PathGetPropertiesAction enumerates the values for path get properties action.
type PathGetPropertiesAction string

const (
	// GetAccessControl ...
	GetAccessControl PathGetPropertiesAction = "getAccessControl"
	// GetStatus ...
	GetStatus PathGetPropertiesAction = "getStatus"
)

// PossiblePathGetPropertiesActionValues returns an array of possible values for the PathGetPropertiesAction const type.
func PossiblePathGetPropertiesActionValues() []PathGetPropertiesAction {
	return []PathGetPropertiesAction{GetAccessControl, GetStatus}
}

// PathLeaseAction enumerates the values for path lease action.
type PathLeaseAction string

const (
	// Acquire ...
	Acquire PathLeaseAction = "acquire"
	// Break ...
	Break PathLeaseAction = "break"
	// Change ...
	Change PathLeaseAction = "change"
	// Release ...
	Release PathLeaseAction = "release"
	// Renew ...
	Renew PathLeaseAction = "renew"
)

// PossiblePathLeaseActionValues returns an array of possible values for the PathLeaseAction const type.
func PossiblePathLeaseActionValues() []PathLeaseAction {
	return []PathLeaseAction{Acquire, Break, Change, Release, Renew}
}

// PathRenameMode enumerates the values for path rename mode.
type PathRenameMode string

const (
	// Legacy ...
	Legacy PathRenameMode = "legacy"
	// Posix ...
	Posix PathRenameMode = "posix"
)

// PossiblePathRenameModeValues returns an array of possible values for the PathRenameMode const type.
func PossiblePathRenameModeValues() []PathRenameMode {
	return []PathRenameMode{Legacy, Posix}
}

// PathResourceType enumerates the values for path resource type.
type PathResourceType string

const (
	// Directory ...
	Directory PathResourceType = "directory"
	// File ...
	File PathResourceType = "file"
)

// PossiblePathResourceTypeValues returns an array of possible values for the PathResourceType const type.
func PossiblePathResourceTypeValues() []PathResourceType {
	return []PathResourceType{Directory, File}
}

// PathUpdateAction enumerates the values for path update action.
type PathUpdateAction string

const (
	// Append ...
	Append PathUpdateAction = "append"
	// Flush ...
	Flush PathUpdateAction = "flush"
	// SetAccessControl ...
	SetAccessControl PathUpdateAction = "setAccessControl"
	// SetProperties ...
	SetProperties PathUpdateAction = "setProperties"
)

// PossiblePathUpdateActionValues returns an array of possible values for the PathUpdateAction const type.
func PossiblePathUpdateActionValues() []PathUpdateAction {
	return []PathUpdateAction{Append, Flush, SetAccessControl, SetProperties}
}
