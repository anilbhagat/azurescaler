package resourcehealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AvailabilityStateValues enumerates the values for availability state values.
type AvailabilityStateValues string

const (
	// Available ...
	Available AvailabilityStateValues = "Available"
	// Unavailable ...
	Unavailable AvailabilityStateValues = "Unavailable"
	// Unknown ...
	Unknown AvailabilityStateValues = "Unknown"
)

// PossibleAvailabilityStateValuesValues returns an array of possible values for the AvailabilityStateValues const type.
func PossibleAvailabilityStateValuesValues() []AvailabilityStateValues {
	return []AvailabilityStateValues{Available, Unavailable, Unknown}
}

// ReasonChronicityTypes enumerates the values for reason chronicity types.
type ReasonChronicityTypes string

const (
	// Persistent ...
	Persistent ReasonChronicityTypes = "Persistent"
	// Transient ...
	Transient ReasonChronicityTypes = "Transient"
)

// PossibleReasonChronicityTypesValues returns an array of possible values for the ReasonChronicityTypes const type.
func PossibleReasonChronicityTypesValues() []ReasonChronicityTypes {
	return []ReasonChronicityTypes{Persistent, Transient}
}