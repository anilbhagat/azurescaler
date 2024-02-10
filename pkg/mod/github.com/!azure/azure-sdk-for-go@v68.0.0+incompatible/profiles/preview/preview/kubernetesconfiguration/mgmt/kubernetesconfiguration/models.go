//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package kubernetesconfiguration

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/kubernetesconfiguration/mgmt/2022-04-02-preview/kubernetesconfiguration"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AKSIdentityType = original.AKSIdentityType

const (
	AKSIdentityTypeSystemAssigned AKSIdentityType = original.AKSIdentityTypeSystemAssigned
	AKSIdentityTypeUserAssigned   AKSIdentityType = original.AKSIdentityTypeUserAssigned
)

type ComplianceStateType = original.ComplianceStateType

const (
	ComplianceStateTypeCompliant    ComplianceStateType = original.ComplianceStateTypeCompliant
	ComplianceStateTypeFailed       ComplianceStateType = original.ComplianceStateTypeFailed
	ComplianceStateTypeInstalled    ComplianceStateType = original.ComplianceStateTypeInstalled
	ComplianceStateTypeNoncompliant ComplianceStateType = original.ComplianceStateTypeNoncompliant
	ComplianceStateTypePending      ComplianceStateType = original.ComplianceStateTypePending
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type FluxComplianceState = original.FluxComplianceState

const (
	FluxComplianceStateCompliant    FluxComplianceState = original.FluxComplianceStateCompliant
	FluxComplianceStateNonCompliant FluxComplianceState = original.FluxComplianceStateNonCompliant
	FluxComplianceStatePending      FluxComplianceState = original.FluxComplianceStatePending
	FluxComplianceStateSuspended    FluxComplianceState = original.FluxComplianceStateSuspended
	FluxComplianceStateUnknown      FluxComplianceState = original.FluxComplianceStateUnknown
)

type KustomizationValidationType = original.KustomizationValidationType

const (
	KustomizationValidationTypeClient KustomizationValidationType = original.KustomizationValidationTypeClient
	KustomizationValidationTypeNone   KustomizationValidationType = original.KustomizationValidationTypeNone
	KustomizationValidationTypeServer KustomizationValidationType = original.KustomizationValidationTypeServer
)

type LevelType = original.LevelType

const (
	LevelTypeError       LevelType = original.LevelTypeError
	LevelTypeInformation LevelType = original.LevelTypeInformation
	LevelTypeWarning     LevelType = original.LevelTypeWarning
)

type MessageLevelType = original.MessageLevelType

const (
	MessageLevelTypeError       MessageLevelType = original.MessageLevelTypeError
	MessageLevelTypeInformation MessageLevelType = original.MessageLevelTypeInformation
	MessageLevelTypeWarning     MessageLevelType = original.MessageLevelTypeWarning
)

type OperatorScopeType = original.OperatorScopeType

const (
	OperatorScopeTypeCluster   OperatorScopeType = original.OperatorScopeTypeCluster
	OperatorScopeTypeNamespace OperatorScopeType = original.OperatorScopeTypeNamespace
)

type OperatorType = original.OperatorType

const (
	OperatorTypeFlux OperatorType = original.OperatorTypeFlux
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusApproved
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusPending
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusRejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type ProvisioningStateType = original.ProvisioningStateType

const (
	ProvisioningStateTypeAccepted  ProvisioningStateType = original.ProvisioningStateTypeAccepted
	ProvisioningStateTypeDeleting  ProvisioningStateType = original.ProvisioningStateTypeDeleting
	ProvisioningStateTypeFailed    ProvisioningStateType = original.ProvisioningStateTypeFailed
	ProvisioningStateTypeRunning   ProvisioningStateType = original.ProvisioningStateTypeRunning
	ProvisioningStateTypeSucceeded ProvisioningStateType = original.ProvisioningStateTypeSucceeded
)

type PublicNetworkAccessType = original.PublicNetworkAccessType

const (
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = original.PublicNetworkAccessTypeDisabled
	PublicNetworkAccessTypeEnabled  PublicNetworkAccessType = original.PublicNetworkAccessTypeEnabled
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
)

type ScopeType = original.ScopeType

const (
	ScopeTypeCluster   ScopeType = original.ScopeTypeCluster
	ScopeTypeNamespace ScopeType = original.ScopeTypeNamespace
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierFree     SkuTier = original.SkuTierFree
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type SourceKindType = original.SourceKindType

const (
	SourceKindTypeBucket        SourceKindType = original.SourceKindTypeBucket
	SourceKindTypeGitRepository SourceKindType = original.SourceKindTypeGitRepository
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BucketDefinition = original.BucketDefinition
type BucketPatchDefinition = original.BucketPatchDefinition
type ComplianceStatus = original.ComplianceStatus
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type Extension = original.Extension
type ExtensionProperties = original.ExtensionProperties
type ExtensionPropertiesAksAssignedIdentity = original.ExtensionPropertiesAksAssignedIdentity
type ExtensionStatus = original.ExtensionStatus
type ExtensionsClient = original.ExtensionsClient
type ExtensionsCreateFuture = original.ExtensionsCreateFuture
type ExtensionsDeleteFuture = original.ExtensionsDeleteFuture
type ExtensionsList = original.ExtensionsList
type ExtensionsListIterator = original.ExtensionsListIterator
type ExtensionsListPage = original.ExtensionsListPage
type ExtensionsUpdateFuture = original.ExtensionsUpdateFuture
type FluxConfigOperationStatusClient = original.FluxConfigOperationStatusClient
type FluxConfiguration = original.FluxConfiguration
type FluxConfigurationPatch = original.FluxConfigurationPatch
type FluxConfigurationPatchProperties = original.FluxConfigurationPatchProperties
type FluxConfigurationProperties = original.FluxConfigurationProperties
type FluxConfigurationsClient = original.FluxConfigurationsClient
type FluxConfigurationsCreateOrUpdateFuture = original.FluxConfigurationsCreateOrUpdateFuture
type FluxConfigurationsDeleteFuture = original.FluxConfigurationsDeleteFuture
type FluxConfigurationsList = original.FluxConfigurationsList
type FluxConfigurationsListIterator = original.FluxConfigurationsListIterator
type FluxConfigurationsListPage = original.FluxConfigurationsListPage
type FluxConfigurationsUpdateFuture = original.FluxConfigurationsUpdateFuture
type GitRepositoryDefinition = original.GitRepositoryDefinition
type GitRepositoryPatchDefinition = original.GitRepositoryPatchDefinition
type HelmOperatorProperties = original.HelmOperatorProperties
type HelmReleasePropertiesDefinition = original.HelmReleasePropertiesDefinition
type Identity = original.Identity
type KustomizationDefinition = original.KustomizationDefinition
type KustomizationPatchDefinition = original.KustomizationPatchDefinition
type ObjectReferenceDefinition = original.ObjectReferenceDefinition
type ObjectStatusConditionDefinition = original.ObjectStatusConditionDefinition
type ObjectStatusDefinition = original.ObjectStatusDefinition
type OperationStatusClient = original.OperationStatusClient
type OperationStatusList = original.OperationStatusList
type OperationStatusListIterator = original.OperationStatusListIterator
type OperationStatusListPage = original.OperationStatusListPage
type OperationStatusResult = original.OperationStatusResult
type OperationsClient = original.OperationsClient
type PatchExtension = original.PatchExtension
type PatchExtensionProperties = original.PatchExtensionProperties
type Plan = original.Plan
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkScope = original.PrivateLinkScope
type PrivateLinkScopeListResult = original.PrivateLinkScopeListResult
type PrivateLinkScopeListResultIterator = original.PrivateLinkScopeListResultIterator
type PrivateLinkScopeListResultPage = original.PrivateLinkScopeListResultPage
type PrivateLinkScopeProperties = original.PrivateLinkScopeProperties
type PrivateLinkScopesClient = original.PrivateLinkScopesClient
type PrivateLinkScopesDeleteFuture = original.PrivateLinkScopesDeleteFuture
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type RepositoryRefDefinition = original.RepositoryRefDefinition
type Resource = original.Resource
type ResourceModelWithAllowedPropertySet = original.ResourceModelWithAllowedPropertySet
type ResourceModelWithAllowedPropertySetIdentity = original.ResourceModelWithAllowedPropertySetIdentity
type ResourceModelWithAllowedPropertySetPlan = original.ResourceModelWithAllowedPropertySetPlan
type ResourceModelWithAllowedPropertySetSku = original.ResourceModelWithAllowedPropertySetSku
type ResourceProviderOperation = original.ResourceProviderOperation
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceProviderOperationListIterator = original.ResourceProviderOperationListIterator
type ResourceProviderOperationListPage = original.ResourceProviderOperationListPage
type Scope = original.Scope
type ScopeCluster = original.ScopeCluster
type ScopeNamespace = original.ScopeNamespace
type Sku = original.Sku
type SourceControlConfiguration = original.SourceControlConfiguration
type SourceControlConfigurationList = original.SourceControlConfigurationList
type SourceControlConfigurationListIterator = original.SourceControlConfigurationListIterator
type SourceControlConfigurationListPage = original.SourceControlConfigurationListPage
type SourceControlConfigurationProperties = original.SourceControlConfigurationProperties
type SourceControlConfigurationsClient = original.SourceControlConfigurationsClient
type SourceControlConfigurationsDeleteFuture = original.SourceControlConfigurationsDeleteFuture
type SystemData = original.SystemData
type TagsResource = original.TagsResource
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClient(subscriptionID)
}
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExtensionsListIterator(page ExtensionsListPage) ExtensionsListIterator {
	return original.NewExtensionsListIterator(page)
}
func NewExtensionsListPage(cur ExtensionsList, getNextPage func(context.Context, ExtensionsList) (ExtensionsList, error)) ExtensionsListPage {
	return original.NewExtensionsListPage(cur, getNextPage)
}
func NewFluxConfigOperationStatusClient(subscriptionID string) FluxConfigOperationStatusClient {
	return original.NewFluxConfigOperationStatusClient(subscriptionID)
}
func NewFluxConfigOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) FluxConfigOperationStatusClient {
	return original.NewFluxConfigOperationStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewFluxConfigurationsClient(subscriptionID string) FluxConfigurationsClient {
	return original.NewFluxConfigurationsClient(subscriptionID)
}
func NewFluxConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) FluxConfigurationsClient {
	return original.NewFluxConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFluxConfigurationsListIterator(page FluxConfigurationsListPage) FluxConfigurationsListIterator {
	return original.NewFluxConfigurationsListIterator(page)
}
func NewFluxConfigurationsListPage(cur FluxConfigurationsList, getNextPage func(context.Context, FluxConfigurationsList) (FluxConfigurationsList, error)) FluxConfigurationsListPage {
	return original.NewFluxConfigurationsListPage(cur, getNextPage)
}
func NewOperationStatusClient(subscriptionID string) OperationStatusClient {
	return original.NewOperationStatusClient(subscriptionID)
}
func NewOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusClient {
	return original.NewOperationStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationStatusListIterator(page OperationStatusListPage) OperationStatusListIterator {
	return original.NewOperationStatusListIterator(page)
}
func NewOperationStatusListPage(cur OperationStatusList, getNextPage func(context.Context, OperationStatusList) (OperationStatusList, error)) OperationStatusListPage {
	return original.NewOperationStatusListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkScopeListResultIterator(page PrivateLinkScopeListResultPage) PrivateLinkScopeListResultIterator {
	return original.NewPrivateLinkScopeListResultIterator(page)
}
func NewPrivateLinkScopeListResultPage(cur PrivateLinkScopeListResult, getNextPage func(context.Context, PrivateLinkScopeListResult) (PrivateLinkScopeListResult, error)) PrivateLinkScopeListResultPage {
	return original.NewPrivateLinkScopeListResultPage(cur, getNextPage)
}
func NewPrivateLinkScopesClient(subscriptionID string) PrivateLinkScopesClient {
	return original.NewPrivateLinkScopesClient(subscriptionID)
}
func NewPrivateLinkScopesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkScopesClient {
	return original.NewPrivateLinkScopesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return original.NewResourceProviderOperationListIterator(page)
}
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return original.NewResourceProviderOperationListPage(cur, getNextPage)
}
func NewSourceControlConfigurationListIterator(page SourceControlConfigurationListPage) SourceControlConfigurationListIterator {
	return original.NewSourceControlConfigurationListIterator(page)
}
func NewSourceControlConfigurationListPage(cur SourceControlConfigurationList, getNextPage func(context.Context, SourceControlConfigurationList) (SourceControlConfigurationList, error)) SourceControlConfigurationListPage {
	return original.NewSourceControlConfigurationListPage(cur, getNextPage)
}
func NewSourceControlConfigurationsClient(subscriptionID string) SourceControlConfigurationsClient {
	return original.NewSourceControlConfigurationsClient(subscriptionID)
}
func NewSourceControlConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) SourceControlConfigurationsClient {
	return original.NewSourceControlConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAKSIdentityTypeValues() []AKSIdentityType {
	return original.PossibleAKSIdentityTypeValues()
}
func PossibleComplianceStateTypeValues() []ComplianceStateType {
	return original.PossibleComplianceStateTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleFluxComplianceStateValues() []FluxComplianceState {
	return original.PossibleFluxComplianceStateValues()
}
func PossibleKustomizationValidationTypeValues() []KustomizationValidationType {
	return original.PossibleKustomizationValidationTypeValues()
}
func PossibleLevelTypeValues() []LevelType {
	return original.PossibleLevelTypeValues()
}
func PossibleMessageLevelTypeValues() []MessageLevelType {
	return original.PossibleMessageLevelTypeValues()
}
func PossibleOperatorScopeTypeValues() []OperatorScopeType {
	return original.PossibleOperatorScopeTypeValues()
}
func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return original.PossibleProvisioningStateTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return original.PossiblePublicNetworkAccessTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleScopeTypeValues() []ScopeType {
	return original.PossibleScopeTypeValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSourceKindTypeValues() []SourceKindType {
	return original.PossibleSourceKindTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
