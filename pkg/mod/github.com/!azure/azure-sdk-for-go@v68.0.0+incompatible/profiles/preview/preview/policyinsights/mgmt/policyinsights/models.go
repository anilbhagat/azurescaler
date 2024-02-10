//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package policyinsights

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2020-07-01-preview/policyinsights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type FieldRestrictionResult = original.FieldRestrictionResult

const (
	Deny     FieldRestrictionResult = original.Deny
	Removed  FieldRestrictionResult = original.Removed
	Required FieldRestrictionResult = original.Required
)

type PolicyStatesResource = original.PolicyStatesResource

const (
	Default PolicyStatesResource = original.Default
	Latest  PolicyStatesResource = original.Latest
)

type ResourceDiscoveryMode = original.ResourceDiscoveryMode

const (
	ExistingNonCompliant ResourceDiscoveryMode = original.ExistingNonCompliant
	ReEvaluateCompliance ResourceDiscoveryMode = original.ReEvaluateCompliance
)

type BaseClient = original.BaseClient
type CheckRestrictionsRequest = original.CheckRestrictionsRequest
type CheckRestrictionsResourceDetails = original.CheckRestrictionsResourceDetails
type CheckRestrictionsResult = original.CheckRestrictionsResult
type CheckRestrictionsResultContentEvaluationResult = original.CheckRestrictionsResultContentEvaluationResult
type ComplianceDetail = original.ComplianceDetail
type ComponentEventDetails = original.ComponentEventDetails
type ComponentStateDetails = original.ComponentStateDetails
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type ExpressionEvaluationDetails = original.ExpressionEvaluationDetails
type FieldRestriction = original.FieldRestriction
type FieldRestrictions = original.FieldRestrictions
type IfNotExistsEvaluationDetails = original.IfNotExistsEvaluationDetails
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsListResults = original.OperationsListResults
type PendingField = original.PendingField
type PolicyAssignmentSummary = original.PolicyAssignmentSummary
type PolicyDefinitionSummary = original.PolicyDefinitionSummary
type PolicyDetails = original.PolicyDetails
type PolicyEvaluationDetails = original.PolicyEvaluationDetails
type PolicyEvaluationResult = original.PolicyEvaluationResult
type PolicyEvent = original.PolicyEvent
type PolicyEventsClient = original.PolicyEventsClient
type PolicyEventsQueryResults = original.PolicyEventsQueryResults
type PolicyEventsQueryResultsIterator = original.PolicyEventsQueryResultsIterator
type PolicyEventsQueryResultsPage = original.PolicyEventsQueryResultsPage
type PolicyGroupSummary = original.PolicyGroupSummary
type PolicyMetadata = original.PolicyMetadata
type PolicyMetadataClient = original.PolicyMetadataClient
type PolicyMetadataCollection = original.PolicyMetadataCollection
type PolicyMetadataCollectionIterator = original.PolicyMetadataCollectionIterator
type PolicyMetadataCollectionPage = original.PolicyMetadataCollectionPage
type PolicyMetadataProperties = original.PolicyMetadataProperties
type PolicyMetadataSlimProperties = original.PolicyMetadataSlimProperties
type PolicyReference = original.PolicyReference
type PolicyRestrictionsClient = original.PolicyRestrictionsClient
type PolicyState = original.PolicyState
type PolicyStatesClient = original.PolicyStatesClient
type PolicyStatesQueryResults = original.PolicyStatesQueryResults
type PolicyStatesQueryResultsIterator = original.PolicyStatesQueryResultsIterator
type PolicyStatesQueryResultsPage = original.PolicyStatesQueryResultsPage
type PolicyStatesTriggerResourceGroupEvaluationFuture = original.PolicyStatesTriggerResourceGroupEvaluationFuture
type PolicyStatesTriggerSubscriptionEvaluationFuture = original.PolicyStatesTriggerSubscriptionEvaluationFuture
type PolicyTrackedResource = original.PolicyTrackedResource
type PolicyTrackedResourcesClient = original.PolicyTrackedResourcesClient
type PolicyTrackedResourcesQueryResults = original.PolicyTrackedResourcesQueryResults
type PolicyTrackedResourcesQueryResultsIterator = original.PolicyTrackedResourcesQueryResultsIterator
type PolicyTrackedResourcesQueryResultsPage = original.PolicyTrackedResourcesQueryResultsPage
type QueryFailure = original.QueryFailure
type QueryFailureError = original.QueryFailureError
type Remediation = original.Remediation
type RemediationDeployment = original.RemediationDeployment
type RemediationDeploymentSummary = original.RemediationDeploymentSummary
type RemediationDeploymentsListResult = original.RemediationDeploymentsListResult
type RemediationDeploymentsListResultIterator = original.RemediationDeploymentsListResultIterator
type RemediationDeploymentsListResultPage = original.RemediationDeploymentsListResultPage
type RemediationFilters = original.RemediationFilters
type RemediationListResult = original.RemediationListResult
type RemediationListResultIterator = original.RemediationListResultIterator
type RemediationListResultPage = original.RemediationListResultPage
type RemediationProperties = original.RemediationProperties
type RemediationsClient = original.RemediationsClient
type SlimPolicyMetadata = original.SlimPolicyMetadata
type SummarizeResults = original.SummarizeResults
type Summary = original.Summary
type SummaryResults = original.SummaryResults
type TrackedResourceModificationDetails = original.TrackedResourceModificationDetails
type TypedErrorInfo = original.TypedErrorInfo

func New(subscriptionID2 string) BaseClient {
	return original.New(subscriptionID2)
}
func NewOperationsClient(subscriptionID2 string) OperationsClient {
	return original.NewOperationsClient(subscriptionID2)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID2 string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID2)
}
func NewPolicyEventsClient(subscriptionID2 string) PolicyEventsClient {
	return original.NewPolicyEventsClient(subscriptionID2)
}
func NewPolicyEventsClientWithBaseURI(baseURI string, subscriptionID2 string) PolicyEventsClient {
	return original.NewPolicyEventsClientWithBaseURI(baseURI, subscriptionID2)
}
func NewPolicyEventsQueryResultsIterator(page PolicyEventsQueryResultsPage) PolicyEventsQueryResultsIterator {
	return original.NewPolicyEventsQueryResultsIterator(page)
}
func NewPolicyEventsQueryResultsPage(cur PolicyEventsQueryResults, getNextPage func(context.Context, PolicyEventsQueryResults) (PolicyEventsQueryResults, error)) PolicyEventsQueryResultsPage {
	return original.NewPolicyEventsQueryResultsPage(cur, getNextPage)
}
func NewPolicyMetadataClient(subscriptionID2 string) PolicyMetadataClient {
	return original.NewPolicyMetadataClient(subscriptionID2)
}
func NewPolicyMetadataClientWithBaseURI(baseURI string, subscriptionID2 string) PolicyMetadataClient {
	return original.NewPolicyMetadataClientWithBaseURI(baseURI, subscriptionID2)
}
func NewPolicyMetadataCollectionIterator(page PolicyMetadataCollectionPage) PolicyMetadataCollectionIterator {
	return original.NewPolicyMetadataCollectionIterator(page)
}
func NewPolicyMetadataCollectionPage(cur PolicyMetadataCollection, getNextPage func(context.Context, PolicyMetadataCollection) (PolicyMetadataCollection, error)) PolicyMetadataCollectionPage {
	return original.NewPolicyMetadataCollectionPage(cur, getNextPage)
}
func NewPolicyRestrictionsClient(subscriptionID2 string) PolicyRestrictionsClient {
	return original.NewPolicyRestrictionsClient(subscriptionID2)
}
func NewPolicyRestrictionsClientWithBaseURI(baseURI string, subscriptionID2 string) PolicyRestrictionsClient {
	return original.NewPolicyRestrictionsClientWithBaseURI(baseURI, subscriptionID2)
}
func NewPolicyStatesClient(subscriptionID2 string) PolicyStatesClient {
	return original.NewPolicyStatesClient(subscriptionID2)
}
func NewPolicyStatesClientWithBaseURI(baseURI string, subscriptionID2 string) PolicyStatesClient {
	return original.NewPolicyStatesClientWithBaseURI(baseURI, subscriptionID2)
}
func NewPolicyStatesQueryResultsIterator(page PolicyStatesQueryResultsPage) PolicyStatesQueryResultsIterator {
	return original.NewPolicyStatesQueryResultsIterator(page)
}
func NewPolicyStatesQueryResultsPage(cur PolicyStatesQueryResults, getNextPage func(context.Context, PolicyStatesQueryResults) (PolicyStatesQueryResults, error)) PolicyStatesQueryResultsPage {
	return original.NewPolicyStatesQueryResultsPage(cur, getNextPage)
}
func NewPolicyTrackedResourcesClient(subscriptionID2 string) PolicyTrackedResourcesClient {
	return original.NewPolicyTrackedResourcesClient(subscriptionID2)
}
func NewPolicyTrackedResourcesClientWithBaseURI(baseURI string, subscriptionID2 string) PolicyTrackedResourcesClient {
	return original.NewPolicyTrackedResourcesClientWithBaseURI(baseURI, subscriptionID2)
}
func NewPolicyTrackedResourcesQueryResultsIterator(page PolicyTrackedResourcesQueryResultsPage) PolicyTrackedResourcesQueryResultsIterator {
	return original.NewPolicyTrackedResourcesQueryResultsIterator(page)
}
func NewPolicyTrackedResourcesQueryResultsPage(cur PolicyTrackedResourcesQueryResults, getNextPage func(context.Context, PolicyTrackedResourcesQueryResults) (PolicyTrackedResourcesQueryResults, error)) PolicyTrackedResourcesQueryResultsPage {
	return original.NewPolicyTrackedResourcesQueryResultsPage(cur, getNextPage)
}
func NewRemediationDeploymentsListResultIterator(page RemediationDeploymentsListResultPage) RemediationDeploymentsListResultIterator {
	return original.NewRemediationDeploymentsListResultIterator(page)
}
func NewRemediationDeploymentsListResultPage(cur RemediationDeploymentsListResult, getNextPage func(context.Context, RemediationDeploymentsListResult) (RemediationDeploymentsListResult, error)) RemediationDeploymentsListResultPage {
	return original.NewRemediationDeploymentsListResultPage(cur, getNextPage)
}
func NewRemediationListResultIterator(page RemediationListResultPage) RemediationListResultIterator {
	return original.NewRemediationListResultIterator(page)
}
func NewRemediationListResultPage(cur RemediationListResult, getNextPage func(context.Context, RemediationListResult) (RemediationListResult, error)) RemediationListResultPage {
	return original.NewRemediationListResultPage(cur, getNextPage)
}
func NewRemediationsClient(subscriptionID2 string) RemediationsClient {
	return original.NewRemediationsClient(subscriptionID2)
}
func NewRemediationsClientWithBaseURI(baseURI string, subscriptionID2 string) RemediationsClient {
	return original.NewRemediationsClientWithBaseURI(baseURI, subscriptionID2)
}
func NewWithBaseURI(baseURI string, subscriptionID2 string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID2)
}
func PossibleFieldRestrictionResultValues() []FieldRestrictionResult {
	return original.PossibleFieldRestrictionResultValues()
}
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return original.PossiblePolicyStatesResourceValues()
}
func PossibleResourceDiscoveryModeValues() []ResourceDiscoveryMode {
	return original.PossibleResourceDiscoveryModeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}