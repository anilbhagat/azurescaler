package authoring

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// OperationStatusType enumerates the values for operation status type.
type OperationStatusType string

const (
	// Failed ...
	Failed OperationStatusType = "Failed"
	// FAILED ...
	FAILED OperationStatusType = "FAILED"
	// Success ...
	Success OperationStatusType = "Success"
)

// PossibleOperationStatusTypeValues returns an array of possible values for the OperationStatusType const type.
func PossibleOperationStatusTypeValues() []OperationStatusType {
	return []OperationStatusType{Failed, FAILED, Success}
}

// ReadableType enumerates the values for readable type.
type ReadableType string

const (
	// ReadableTypeChildEntityExtractor ...
	ReadableTypeChildEntityExtractor ReadableType = "Child Entity Extractor"
	// ReadableTypeClosedListEntityExtractor ...
	ReadableTypeClosedListEntityExtractor ReadableType = "Closed List Entity Extractor"
	// ReadableTypeCompositeEntityExtractor ...
	ReadableTypeCompositeEntityExtractor ReadableType = "Composite Entity Extractor"
	// ReadableTypeEntityExtractor ...
	ReadableTypeEntityExtractor ReadableType = "Entity Extractor"
	// ReadableTypeHierarchicalChildEntityExtractor ...
	ReadableTypeHierarchicalChildEntityExtractor ReadableType = "Hierarchical Child Entity Extractor"
	// ReadableTypeHierarchicalEntityExtractor ...
	ReadableTypeHierarchicalEntityExtractor ReadableType = "Hierarchical Entity Extractor"
	// ReadableTypeIntentClassifier ...
	ReadableTypeIntentClassifier ReadableType = "Intent Classifier"
	// ReadableTypeListEntityExtractor ...
	ReadableTypeListEntityExtractor ReadableType = "List Entity Extractor"
	// ReadableTypePatternAnyEntityExtractor ...
	ReadableTypePatternAnyEntityExtractor ReadableType = "Pattern.Any Entity Extractor"
	// ReadableTypePrebuiltEntityExtractor ...
	ReadableTypePrebuiltEntityExtractor ReadableType = "Prebuilt Entity Extractor"
	// ReadableTypeRegexEntityExtractor ...
	ReadableTypeRegexEntityExtractor ReadableType = "Regex Entity Extractor"
)

// PossibleReadableTypeValues returns an array of possible values for the ReadableType const type.
func PossibleReadableTypeValues() []ReadableType {
	return []ReadableType{ReadableTypeChildEntityExtractor, ReadableTypeClosedListEntityExtractor, ReadableTypeCompositeEntityExtractor, ReadableTypeEntityExtractor, ReadableTypeHierarchicalChildEntityExtractor, ReadableTypeHierarchicalEntityExtractor, ReadableTypeIntentClassifier, ReadableTypeListEntityExtractor, ReadableTypePatternAnyEntityExtractor, ReadableTypePrebuiltEntityExtractor, ReadableTypeRegexEntityExtractor}
}

// ReadableType1 enumerates the values for readable type 1.
type ReadableType1 string

const (
	// ReadableType1ChildEntityExtractor ...
	ReadableType1ChildEntityExtractor ReadableType1 = "Child Entity Extractor"
	// ReadableType1ClosedListEntityExtractor ...
	ReadableType1ClosedListEntityExtractor ReadableType1 = "Closed List Entity Extractor"
	// ReadableType1CompositeEntityExtractor ...
	ReadableType1CompositeEntityExtractor ReadableType1 = "Composite Entity Extractor"
	// ReadableType1EntityExtractor ...
	ReadableType1EntityExtractor ReadableType1 = "Entity Extractor"
	// ReadableType1HierarchicalChildEntityExtractor ...
	ReadableType1HierarchicalChildEntityExtractor ReadableType1 = "Hierarchical Child Entity Extractor"
	// ReadableType1HierarchicalEntityExtractor ...
	ReadableType1HierarchicalEntityExtractor ReadableType1 = "Hierarchical Entity Extractor"
	// ReadableType1IntentClassifier ...
	ReadableType1IntentClassifier ReadableType1 = "Intent Classifier"
	// ReadableType1ListEntityExtractor ...
	ReadableType1ListEntityExtractor ReadableType1 = "List Entity Extractor"
	// ReadableType1PatternAnyEntityExtractor ...
	ReadableType1PatternAnyEntityExtractor ReadableType1 = "Pattern.Any Entity Extractor"
	// ReadableType1PrebuiltEntityExtractor ...
	ReadableType1PrebuiltEntityExtractor ReadableType1 = "Prebuilt Entity Extractor"
	// ReadableType1RegexEntityExtractor ...
	ReadableType1RegexEntityExtractor ReadableType1 = "Regex Entity Extractor"
)

// PossibleReadableType1Values returns an array of possible values for the ReadableType1 const type.
func PossibleReadableType1Values() []ReadableType1 {
	return []ReadableType1{ReadableType1ChildEntityExtractor, ReadableType1ClosedListEntityExtractor, ReadableType1CompositeEntityExtractor, ReadableType1EntityExtractor, ReadableType1HierarchicalChildEntityExtractor, ReadableType1HierarchicalEntityExtractor, ReadableType1IntentClassifier, ReadableType1ListEntityExtractor, ReadableType1PatternAnyEntityExtractor, ReadableType1PrebuiltEntityExtractor, ReadableType1RegexEntityExtractor}
}

// ReadableType10 enumerates the values for readable type 10.
type ReadableType10 string

const (
	// ReadableType10ChildEntityExtractor ...
	ReadableType10ChildEntityExtractor ReadableType10 = "Child Entity Extractor"
	// ReadableType10ClosedListEntityExtractor ...
	ReadableType10ClosedListEntityExtractor ReadableType10 = "Closed List Entity Extractor"
	// ReadableType10CompositeEntityExtractor ...
	ReadableType10CompositeEntityExtractor ReadableType10 = "Composite Entity Extractor"
	// ReadableType10EntityExtractor ...
	ReadableType10EntityExtractor ReadableType10 = "Entity Extractor"
	// ReadableType10HierarchicalChildEntityExtractor ...
	ReadableType10HierarchicalChildEntityExtractor ReadableType10 = "Hierarchical Child Entity Extractor"
	// ReadableType10HierarchicalEntityExtractor ...
	ReadableType10HierarchicalEntityExtractor ReadableType10 = "Hierarchical Entity Extractor"
	// ReadableType10IntentClassifier ...
	ReadableType10IntentClassifier ReadableType10 = "Intent Classifier"
	// ReadableType10ListEntityExtractor ...
	ReadableType10ListEntityExtractor ReadableType10 = "List Entity Extractor"
	// ReadableType10PatternAnyEntityExtractor ...
	ReadableType10PatternAnyEntityExtractor ReadableType10 = "Pattern.Any Entity Extractor"
	// ReadableType10PrebuiltEntityExtractor ...
	ReadableType10PrebuiltEntityExtractor ReadableType10 = "Prebuilt Entity Extractor"
	// ReadableType10RegexEntityExtractor ...
	ReadableType10RegexEntityExtractor ReadableType10 = "Regex Entity Extractor"
)

// PossibleReadableType10Values returns an array of possible values for the ReadableType10 const type.
func PossibleReadableType10Values() []ReadableType10 {
	return []ReadableType10{ReadableType10ChildEntityExtractor, ReadableType10ClosedListEntityExtractor, ReadableType10CompositeEntityExtractor, ReadableType10EntityExtractor, ReadableType10HierarchicalChildEntityExtractor, ReadableType10HierarchicalEntityExtractor, ReadableType10IntentClassifier, ReadableType10ListEntityExtractor, ReadableType10PatternAnyEntityExtractor, ReadableType10PrebuiltEntityExtractor, ReadableType10RegexEntityExtractor}
}

// ReadableType11 enumerates the values for readable type 11.
type ReadableType11 string

const (
	// ReadableType11ChildEntityExtractor ...
	ReadableType11ChildEntityExtractor ReadableType11 = "Child Entity Extractor"
	// ReadableType11ClosedListEntityExtractor ...
	ReadableType11ClosedListEntityExtractor ReadableType11 = "Closed List Entity Extractor"
	// ReadableType11CompositeEntityExtractor ...
	ReadableType11CompositeEntityExtractor ReadableType11 = "Composite Entity Extractor"
	// ReadableType11EntityExtractor ...
	ReadableType11EntityExtractor ReadableType11 = "Entity Extractor"
	// ReadableType11HierarchicalChildEntityExtractor ...
	ReadableType11HierarchicalChildEntityExtractor ReadableType11 = "Hierarchical Child Entity Extractor"
	// ReadableType11HierarchicalEntityExtractor ...
	ReadableType11HierarchicalEntityExtractor ReadableType11 = "Hierarchical Entity Extractor"
	// ReadableType11IntentClassifier ...
	ReadableType11IntentClassifier ReadableType11 = "Intent Classifier"
	// ReadableType11ListEntityExtractor ...
	ReadableType11ListEntityExtractor ReadableType11 = "List Entity Extractor"
	// ReadableType11PatternAnyEntityExtractor ...
	ReadableType11PatternAnyEntityExtractor ReadableType11 = "Pattern.Any Entity Extractor"
	// ReadableType11PrebuiltEntityExtractor ...
	ReadableType11PrebuiltEntityExtractor ReadableType11 = "Prebuilt Entity Extractor"
	// ReadableType11RegexEntityExtractor ...
	ReadableType11RegexEntityExtractor ReadableType11 = "Regex Entity Extractor"
)

// PossibleReadableType11Values returns an array of possible values for the ReadableType11 const type.
func PossibleReadableType11Values() []ReadableType11 {
	return []ReadableType11{ReadableType11ChildEntityExtractor, ReadableType11ClosedListEntityExtractor, ReadableType11CompositeEntityExtractor, ReadableType11EntityExtractor, ReadableType11HierarchicalChildEntityExtractor, ReadableType11HierarchicalEntityExtractor, ReadableType11IntentClassifier, ReadableType11ListEntityExtractor, ReadableType11PatternAnyEntityExtractor, ReadableType11PrebuiltEntityExtractor, ReadableType11RegexEntityExtractor}
}

// ReadableType2 enumerates the values for readable type 2.
type ReadableType2 string

const (
	// ReadableType2ChildEntityExtractor ...
	ReadableType2ChildEntityExtractor ReadableType2 = "Child Entity Extractor"
	// ReadableType2ClosedListEntityExtractor ...
	ReadableType2ClosedListEntityExtractor ReadableType2 = "Closed List Entity Extractor"
	// ReadableType2CompositeEntityExtractor ...
	ReadableType2CompositeEntityExtractor ReadableType2 = "Composite Entity Extractor"
	// ReadableType2EntityExtractor ...
	ReadableType2EntityExtractor ReadableType2 = "Entity Extractor"
	// ReadableType2HierarchicalChildEntityExtractor ...
	ReadableType2HierarchicalChildEntityExtractor ReadableType2 = "Hierarchical Child Entity Extractor"
	// ReadableType2HierarchicalEntityExtractor ...
	ReadableType2HierarchicalEntityExtractor ReadableType2 = "Hierarchical Entity Extractor"
	// ReadableType2IntentClassifier ...
	ReadableType2IntentClassifier ReadableType2 = "Intent Classifier"
	// ReadableType2ListEntityExtractor ...
	ReadableType2ListEntityExtractor ReadableType2 = "List Entity Extractor"
	// ReadableType2PatternAnyEntityExtractor ...
	ReadableType2PatternAnyEntityExtractor ReadableType2 = "Pattern.Any Entity Extractor"
	// ReadableType2PrebuiltEntityExtractor ...
	ReadableType2PrebuiltEntityExtractor ReadableType2 = "Prebuilt Entity Extractor"
	// ReadableType2RegexEntityExtractor ...
	ReadableType2RegexEntityExtractor ReadableType2 = "Regex Entity Extractor"
)

// PossibleReadableType2Values returns an array of possible values for the ReadableType2 const type.
func PossibleReadableType2Values() []ReadableType2 {
	return []ReadableType2{ReadableType2ChildEntityExtractor, ReadableType2ClosedListEntityExtractor, ReadableType2CompositeEntityExtractor, ReadableType2EntityExtractor, ReadableType2HierarchicalChildEntityExtractor, ReadableType2HierarchicalEntityExtractor, ReadableType2IntentClassifier, ReadableType2ListEntityExtractor, ReadableType2PatternAnyEntityExtractor, ReadableType2PrebuiltEntityExtractor, ReadableType2RegexEntityExtractor}
}

// ReadableType3 enumerates the values for readable type 3.
type ReadableType3 string

const (
	// ReadableType3ChildEntityExtractor ...
	ReadableType3ChildEntityExtractor ReadableType3 = "Child Entity Extractor"
	// ReadableType3ClosedListEntityExtractor ...
	ReadableType3ClosedListEntityExtractor ReadableType3 = "Closed List Entity Extractor"
	// ReadableType3CompositeEntityExtractor ...
	ReadableType3CompositeEntityExtractor ReadableType3 = "Composite Entity Extractor"
	// ReadableType3EntityExtractor ...
	ReadableType3EntityExtractor ReadableType3 = "Entity Extractor"
	// ReadableType3HierarchicalChildEntityExtractor ...
	ReadableType3HierarchicalChildEntityExtractor ReadableType3 = "Hierarchical Child Entity Extractor"
	// ReadableType3HierarchicalEntityExtractor ...
	ReadableType3HierarchicalEntityExtractor ReadableType3 = "Hierarchical Entity Extractor"
	// ReadableType3IntentClassifier ...
	ReadableType3IntentClassifier ReadableType3 = "Intent Classifier"
	// ReadableType3ListEntityExtractor ...
	ReadableType3ListEntityExtractor ReadableType3 = "List Entity Extractor"
	// ReadableType3PatternAnyEntityExtractor ...
	ReadableType3PatternAnyEntityExtractor ReadableType3 = "Pattern.Any Entity Extractor"
	// ReadableType3PrebuiltEntityExtractor ...
	ReadableType3PrebuiltEntityExtractor ReadableType3 = "Prebuilt Entity Extractor"
	// ReadableType3RegexEntityExtractor ...
	ReadableType3RegexEntityExtractor ReadableType3 = "Regex Entity Extractor"
)

// PossibleReadableType3Values returns an array of possible values for the ReadableType3 const type.
func PossibleReadableType3Values() []ReadableType3 {
	return []ReadableType3{ReadableType3ChildEntityExtractor, ReadableType3ClosedListEntityExtractor, ReadableType3CompositeEntityExtractor, ReadableType3EntityExtractor, ReadableType3HierarchicalChildEntityExtractor, ReadableType3HierarchicalEntityExtractor, ReadableType3IntentClassifier, ReadableType3ListEntityExtractor, ReadableType3PatternAnyEntityExtractor, ReadableType3PrebuiltEntityExtractor, ReadableType3RegexEntityExtractor}
}

// ReadableType4 enumerates the values for readable type 4.
type ReadableType4 string

const (
	// ReadableType4ChildEntityExtractor ...
	ReadableType4ChildEntityExtractor ReadableType4 = "Child Entity Extractor"
	// ReadableType4ClosedListEntityExtractor ...
	ReadableType4ClosedListEntityExtractor ReadableType4 = "Closed List Entity Extractor"
	// ReadableType4CompositeEntityExtractor ...
	ReadableType4CompositeEntityExtractor ReadableType4 = "Composite Entity Extractor"
	// ReadableType4EntityExtractor ...
	ReadableType4EntityExtractor ReadableType4 = "Entity Extractor"
	// ReadableType4HierarchicalChildEntityExtractor ...
	ReadableType4HierarchicalChildEntityExtractor ReadableType4 = "Hierarchical Child Entity Extractor"
	// ReadableType4HierarchicalEntityExtractor ...
	ReadableType4HierarchicalEntityExtractor ReadableType4 = "Hierarchical Entity Extractor"
	// ReadableType4IntentClassifier ...
	ReadableType4IntentClassifier ReadableType4 = "Intent Classifier"
	// ReadableType4ListEntityExtractor ...
	ReadableType4ListEntityExtractor ReadableType4 = "List Entity Extractor"
	// ReadableType4PatternAnyEntityExtractor ...
	ReadableType4PatternAnyEntityExtractor ReadableType4 = "Pattern.Any Entity Extractor"
	// ReadableType4PrebuiltEntityExtractor ...
	ReadableType4PrebuiltEntityExtractor ReadableType4 = "Prebuilt Entity Extractor"
	// ReadableType4RegexEntityExtractor ...
	ReadableType4RegexEntityExtractor ReadableType4 = "Regex Entity Extractor"
)

// PossibleReadableType4Values returns an array of possible values for the ReadableType4 const type.
func PossibleReadableType4Values() []ReadableType4 {
	return []ReadableType4{ReadableType4ChildEntityExtractor, ReadableType4ClosedListEntityExtractor, ReadableType4CompositeEntityExtractor, ReadableType4EntityExtractor, ReadableType4HierarchicalChildEntityExtractor, ReadableType4HierarchicalEntityExtractor, ReadableType4IntentClassifier, ReadableType4ListEntityExtractor, ReadableType4PatternAnyEntityExtractor, ReadableType4PrebuiltEntityExtractor, ReadableType4RegexEntityExtractor}
}

// ReadableType5 enumerates the values for readable type 5.
type ReadableType5 string

const (
	// ReadableType5ChildEntityExtractor ...
	ReadableType5ChildEntityExtractor ReadableType5 = "Child Entity Extractor"
	// ReadableType5ClosedListEntityExtractor ...
	ReadableType5ClosedListEntityExtractor ReadableType5 = "Closed List Entity Extractor"
	// ReadableType5CompositeEntityExtractor ...
	ReadableType5CompositeEntityExtractor ReadableType5 = "Composite Entity Extractor"
	// ReadableType5EntityExtractor ...
	ReadableType5EntityExtractor ReadableType5 = "Entity Extractor"
	// ReadableType5HierarchicalChildEntityExtractor ...
	ReadableType5HierarchicalChildEntityExtractor ReadableType5 = "Hierarchical Child Entity Extractor"
	// ReadableType5HierarchicalEntityExtractor ...
	ReadableType5HierarchicalEntityExtractor ReadableType5 = "Hierarchical Entity Extractor"
	// ReadableType5IntentClassifier ...
	ReadableType5IntentClassifier ReadableType5 = "Intent Classifier"
	// ReadableType5ListEntityExtractor ...
	ReadableType5ListEntityExtractor ReadableType5 = "List Entity Extractor"
	// ReadableType5PatternAnyEntityExtractor ...
	ReadableType5PatternAnyEntityExtractor ReadableType5 = "Pattern.Any Entity Extractor"
	// ReadableType5PrebuiltEntityExtractor ...
	ReadableType5PrebuiltEntityExtractor ReadableType5 = "Prebuilt Entity Extractor"
	// ReadableType5RegexEntityExtractor ...
	ReadableType5RegexEntityExtractor ReadableType5 = "Regex Entity Extractor"
)

// PossibleReadableType5Values returns an array of possible values for the ReadableType5 const type.
func PossibleReadableType5Values() []ReadableType5 {
	return []ReadableType5{ReadableType5ChildEntityExtractor, ReadableType5ClosedListEntityExtractor, ReadableType5CompositeEntityExtractor, ReadableType5EntityExtractor, ReadableType5HierarchicalChildEntityExtractor, ReadableType5HierarchicalEntityExtractor, ReadableType5IntentClassifier, ReadableType5ListEntityExtractor, ReadableType5PatternAnyEntityExtractor, ReadableType5PrebuiltEntityExtractor, ReadableType5RegexEntityExtractor}
}

// ReadableType6 enumerates the values for readable type 6.
type ReadableType6 string

const (
	// ReadableType6ChildEntityExtractor ...
	ReadableType6ChildEntityExtractor ReadableType6 = "Child Entity Extractor"
	// ReadableType6ClosedListEntityExtractor ...
	ReadableType6ClosedListEntityExtractor ReadableType6 = "Closed List Entity Extractor"
	// ReadableType6CompositeEntityExtractor ...
	ReadableType6CompositeEntityExtractor ReadableType6 = "Composite Entity Extractor"
	// ReadableType6EntityExtractor ...
	ReadableType6EntityExtractor ReadableType6 = "Entity Extractor"
	// ReadableType6HierarchicalChildEntityExtractor ...
	ReadableType6HierarchicalChildEntityExtractor ReadableType6 = "Hierarchical Child Entity Extractor"
	// ReadableType6HierarchicalEntityExtractor ...
	ReadableType6HierarchicalEntityExtractor ReadableType6 = "Hierarchical Entity Extractor"
	// ReadableType6IntentClassifier ...
	ReadableType6IntentClassifier ReadableType6 = "Intent Classifier"
	// ReadableType6ListEntityExtractor ...
	ReadableType6ListEntityExtractor ReadableType6 = "List Entity Extractor"
	// ReadableType6PatternAnyEntityExtractor ...
	ReadableType6PatternAnyEntityExtractor ReadableType6 = "Pattern.Any Entity Extractor"
	// ReadableType6PrebuiltEntityExtractor ...
	ReadableType6PrebuiltEntityExtractor ReadableType6 = "Prebuilt Entity Extractor"
	// ReadableType6RegexEntityExtractor ...
	ReadableType6RegexEntityExtractor ReadableType6 = "Regex Entity Extractor"
)

// PossibleReadableType6Values returns an array of possible values for the ReadableType6 const type.
func PossibleReadableType6Values() []ReadableType6 {
	return []ReadableType6{ReadableType6ChildEntityExtractor, ReadableType6ClosedListEntityExtractor, ReadableType6CompositeEntityExtractor, ReadableType6EntityExtractor, ReadableType6HierarchicalChildEntityExtractor, ReadableType6HierarchicalEntityExtractor, ReadableType6IntentClassifier, ReadableType6ListEntityExtractor, ReadableType6PatternAnyEntityExtractor, ReadableType6PrebuiltEntityExtractor, ReadableType6RegexEntityExtractor}
}

// ReadableType7 enumerates the values for readable type 7.
type ReadableType7 string

const (
	// ReadableType7ChildEntityExtractor ...
	ReadableType7ChildEntityExtractor ReadableType7 = "Child Entity Extractor"
	// ReadableType7ClosedListEntityExtractor ...
	ReadableType7ClosedListEntityExtractor ReadableType7 = "Closed List Entity Extractor"
	// ReadableType7CompositeEntityExtractor ...
	ReadableType7CompositeEntityExtractor ReadableType7 = "Composite Entity Extractor"
	// ReadableType7EntityExtractor ...
	ReadableType7EntityExtractor ReadableType7 = "Entity Extractor"
	// ReadableType7HierarchicalChildEntityExtractor ...
	ReadableType7HierarchicalChildEntityExtractor ReadableType7 = "Hierarchical Child Entity Extractor"
	// ReadableType7HierarchicalEntityExtractor ...
	ReadableType7HierarchicalEntityExtractor ReadableType7 = "Hierarchical Entity Extractor"
	// ReadableType7IntentClassifier ...
	ReadableType7IntentClassifier ReadableType7 = "Intent Classifier"
	// ReadableType7ListEntityExtractor ...
	ReadableType7ListEntityExtractor ReadableType7 = "List Entity Extractor"
	// ReadableType7PatternAnyEntityExtractor ...
	ReadableType7PatternAnyEntityExtractor ReadableType7 = "Pattern.Any Entity Extractor"
	// ReadableType7PrebuiltEntityExtractor ...
	ReadableType7PrebuiltEntityExtractor ReadableType7 = "Prebuilt Entity Extractor"
	// ReadableType7RegexEntityExtractor ...
	ReadableType7RegexEntityExtractor ReadableType7 = "Regex Entity Extractor"
)

// PossibleReadableType7Values returns an array of possible values for the ReadableType7 const type.
func PossibleReadableType7Values() []ReadableType7 {
	return []ReadableType7{ReadableType7ChildEntityExtractor, ReadableType7ClosedListEntityExtractor, ReadableType7CompositeEntityExtractor, ReadableType7EntityExtractor, ReadableType7HierarchicalChildEntityExtractor, ReadableType7HierarchicalEntityExtractor, ReadableType7IntentClassifier, ReadableType7ListEntityExtractor, ReadableType7PatternAnyEntityExtractor, ReadableType7PrebuiltEntityExtractor, ReadableType7RegexEntityExtractor}
}

// ReadableType8 enumerates the values for readable type 8.
type ReadableType8 string

const (
	// ReadableType8ChildEntityExtractor ...
	ReadableType8ChildEntityExtractor ReadableType8 = "Child Entity Extractor"
	// ReadableType8ClosedListEntityExtractor ...
	ReadableType8ClosedListEntityExtractor ReadableType8 = "Closed List Entity Extractor"
	// ReadableType8CompositeEntityExtractor ...
	ReadableType8CompositeEntityExtractor ReadableType8 = "Composite Entity Extractor"
	// ReadableType8EntityExtractor ...
	ReadableType8EntityExtractor ReadableType8 = "Entity Extractor"
	// ReadableType8HierarchicalChildEntityExtractor ...
	ReadableType8HierarchicalChildEntityExtractor ReadableType8 = "Hierarchical Child Entity Extractor"
	// ReadableType8HierarchicalEntityExtractor ...
	ReadableType8HierarchicalEntityExtractor ReadableType8 = "Hierarchical Entity Extractor"
	// ReadableType8IntentClassifier ...
	ReadableType8IntentClassifier ReadableType8 = "Intent Classifier"
	// ReadableType8ListEntityExtractor ...
	ReadableType8ListEntityExtractor ReadableType8 = "List Entity Extractor"
	// ReadableType8PatternAnyEntityExtractor ...
	ReadableType8PatternAnyEntityExtractor ReadableType8 = "Pattern.Any Entity Extractor"
	// ReadableType8PrebuiltEntityExtractor ...
	ReadableType8PrebuiltEntityExtractor ReadableType8 = "Prebuilt Entity Extractor"
	// ReadableType8RegexEntityExtractor ...
	ReadableType8RegexEntityExtractor ReadableType8 = "Regex Entity Extractor"
)

// PossibleReadableType8Values returns an array of possible values for the ReadableType8 const type.
func PossibleReadableType8Values() []ReadableType8 {
	return []ReadableType8{ReadableType8ChildEntityExtractor, ReadableType8ClosedListEntityExtractor, ReadableType8CompositeEntityExtractor, ReadableType8EntityExtractor, ReadableType8HierarchicalChildEntityExtractor, ReadableType8HierarchicalEntityExtractor, ReadableType8IntentClassifier, ReadableType8ListEntityExtractor, ReadableType8PatternAnyEntityExtractor, ReadableType8PrebuiltEntityExtractor, ReadableType8RegexEntityExtractor}
}

// ReadableType9 enumerates the values for readable type 9.
type ReadableType9 string

const (
	// ReadableType9ChildEntityExtractor ...
	ReadableType9ChildEntityExtractor ReadableType9 = "Child Entity Extractor"
	// ReadableType9ClosedListEntityExtractor ...
	ReadableType9ClosedListEntityExtractor ReadableType9 = "Closed List Entity Extractor"
	// ReadableType9CompositeEntityExtractor ...
	ReadableType9CompositeEntityExtractor ReadableType9 = "Composite Entity Extractor"
	// ReadableType9EntityExtractor ...
	ReadableType9EntityExtractor ReadableType9 = "Entity Extractor"
	// ReadableType9HierarchicalChildEntityExtractor ...
	ReadableType9HierarchicalChildEntityExtractor ReadableType9 = "Hierarchical Child Entity Extractor"
	// ReadableType9HierarchicalEntityExtractor ...
	ReadableType9HierarchicalEntityExtractor ReadableType9 = "Hierarchical Entity Extractor"
	// ReadableType9IntentClassifier ...
	ReadableType9IntentClassifier ReadableType9 = "Intent Classifier"
	// ReadableType9ListEntityExtractor ...
	ReadableType9ListEntityExtractor ReadableType9 = "List Entity Extractor"
	// ReadableType9PatternAnyEntityExtractor ...
	ReadableType9PatternAnyEntityExtractor ReadableType9 = "Pattern.Any Entity Extractor"
	// ReadableType9PrebuiltEntityExtractor ...
	ReadableType9PrebuiltEntityExtractor ReadableType9 = "Prebuilt Entity Extractor"
	// ReadableType9RegexEntityExtractor ...
	ReadableType9RegexEntityExtractor ReadableType9 = "Regex Entity Extractor"
)

// PossibleReadableType9Values returns an array of possible values for the ReadableType9 const type.
func PossibleReadableType9Values() []ReadableType9 {
	return []ReadableType9{ReadableType9ChildEntityExtractor, ReadableType9ClosedListEntityExtractor, ReadableType9CompositeEntityExtractor, ReadableType9EntityExtractor, ReadableType9HierarchicalChildEntityExtractor, ReadableType9HierarchicalEntityExtractor, ReadableType9IntentClassifier, ReadableType9ListEntityExtractor, ReadableType9PatternAnyEntityExtractor, ReadableType9PrebuiltEntityExtractor, ReadableType9RegexEntityExtractor}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusFail ...
	StatusFail Status = "Fail"
	// StatusInProgress ...
	StatusInProgress Status = "InProgress"
	// StatusQueued ...
	StatusQueued Status = "Queued"
	// StatusSuccess ...
	StatusSuccess Status = "Success"
	// StatusUpToDate ...
	StatusUpToDate Status = "UpToDate"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusFail, StatusInProgress, StatusQueued, StatusSuccess, StatusUpToDate}
}

// Status1 enumerates the values for status 1.
type Status1 string

const (
	// Status1Fail ...
	Status1Fail Status1 = "Fail"
	// Status1InProgress ...
	Status1InProgress Status1 = "InProgress"
	// Status1Queued ...
	Status1Queued Status1 = "Queued"
	// Status1Success ...
	Status1Success Status1 = "Success"
	// Status1UpToDate ...
	Status1UpToDate Status1 = "UpToDate"
)

// PossibleStatus1Values returns an array of possible values for the Status1 const type.
func PossibleStatus1Values() []Status1 {
	return []Status1{Status1Fail, Status1InProgress, Status1Queued, Status1Success, Status1UpToDate}
}

// TrainingStatus enumerates the values for training status.
type TrainingStatus string

const (
	// InProgress ...
	InProgress TrainingStatus = "InProgress"
	// NeedsTraining ...
	NeedsTraining TrainingStatus = "NeedsTraining"
	// Trained ...
	Trained TrainingStatus = "Trained"
)

// PossibleTrainingStatusValues returns an array of possible values for the TrainingStatus const type.
func PossibleTrainingStatusValues() []TrainingStatus {
	return []TrainingStatus{InProgress, NeedsTraining, Trained}
}