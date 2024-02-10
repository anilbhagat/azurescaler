package costmanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// FormatType enumerates the values for format type.
type FormatType string

const (
	// Csv ...
	Csv FormatType = "Csv"
)

// PossibleFormatTypeValues returns an array of possible values for the FormatType const type.
func PossibleFormatTypeValues() []FormatType {
	return []FormatType{Csv}
}

// GranularityType enumerates the values for granularity type.
type GranularityType string

const (
	// Daily ...
	Daily GranularityType = "Daily"
)

// PossibleGranularityTypeValues returns an array of possible values for the GranularityType const type.
func PossibleGranularityTypeValues() []GranularityType {
	return []GranularityType{Daily}
}

// RecurrenceType enumerates the values for recurrence type.
type RecurrenceType string

const (
	// RecurrenceTypeAnnually ...
	RecurrenceTypeAnnually RecurrenceType = "Annually"
	// RecurrenceTypeDaily ...
	RecurrenceTypeDaily RecurrenceType = "Daily"
	// RecurrenceTypeMonthly ...
	RecurrenceTypeMonthly RecurrenceType = "Monthly"
	// RecurrenceTypeWeekly ...
	RecurrenceTypeWeekly RecurrenceType = "Weekly"
)

// PossibleRecurrenceTypeValues returns an array of possible values for the RecurrenceType const type.
func PossibleRecurrenceTypeValues() []RecurrenceType {
	return []RecurrenceType{RecurrenceTypeAnnually, RecurrenceTypeDaily, RecurrenceTypeMonthly, RecurrenceTypeWeekly}
}

// ReportConfigColumnType enumerates the values for report config column type.
type ReportConfigColumnType string

const (
	// ReportConfigColumnTypeDimension ...
	ReportConfigColumnTypeDimension ReportConfigColumnType = "Dimension"
	// ReportConfigColumnTypeTag ...
	ReportConfigColumnTypeTag ReportConfigColumnType = "Tag"
)

// PossibleReportConfigColumnTypeValues returns an array of possible values for the ReportConfigColumnType const type.
func PossibleReportConfigColumnTypeValues() []ReportConfigColumnType {
	return []ReportConfigColumnType{ReportConfigColumnTypeDimension, ReportConfigColumnTypeTag}
}

// StatusType enumerates the values for status type.
type StatusType string

const (
	// Active ...
	Active StatusType = "Active"
	// Inactive ...
	Inactive StatusType = "Inactive"
)

// PossibleStatusTypeValues returns an array of possible values for the StatusType const type.
func PossibleStatusTypeValues() []StatusType {
	return []StatusType{Active, Inactive}
}

// TimeframeType enumerates the values for timeframe type.
type TimeframeType string

const (
	// Custom ...
	Custom TimeframeType = "Custom"
	// MonthToDate ...
	MonthToDate TimeframeType = "MonthToDate"
	// WeekToDate ...
	WeekToDate TimeframeType = "WeekToDate"
	// YearToDate ...
	YearToDate TimeframeType = "YearToDate"
)

// PossibleTimeframeTypeValues returns an array of possible values for the TimeframeType const type.
func PossibleTimeframeTypeValues() []TimeframeType {
	return []TimeframeType{Custom, MonthToDate, WeekToDate, YearToDate}
}