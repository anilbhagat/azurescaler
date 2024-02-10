package policyinsights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// PolicyStatesResource enumerates the values for policy states resource.
type PolicyStatesResource string

const (
	// Default ...
	Default PolicyStatesResource = "default"
	// Latest ...
	Latest PolicyStatesResource = "latest"
)

// PossiblePolicyStatesResourceValues returns an array of possible values for the PolicyStatesResource const type.
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return []PolicyStatesResource{Default, Latest}
}
