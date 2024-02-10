//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package search

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AdminKeyKind = original.AdminKeyKind

const (
	Primary   AdminKeyKind = original.Primary
	Secondary AdminKeyKind = original.Secondary
)

type HostingMode = original.HostingMode

const (
	Default     HostingMode = original.Default
	HighDensity HostingMode = original.HighDensity
)

type IdentityType = original.IdentityType

const (
	None           IdentityType = original.None
	SystemAssigned IdentityType = original.SystemAssigned
)

type PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatus

const (
	Approved     PrivateLinkServiceConnectionStatus = original.Approved
	Disconnected PrivateLinkServiceConnectionStatus = original.Disconnected
	Pending      PrivateLinkServiceConnectionStatus = original.Pending
	Rejected     PrivateLinkServiceConnectionStatus = original.Rejected
)

type ProvisioningState = original.ProvisioningState

const (
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	Disabled PublicNetworkAccess = original.Disabled
	Enabled  PublicNetworkAccess = original.Enabled
)

type ServiceStatus = original.ServiceStatus

const (
	ServiceStatusDegraded     ServiceStatus = original.ServiceStatusDegraded
	ServiceStatusDeleting     ServiceStatus = original.ServiceStatusDeleting
	ServiceStatusDisabled     ServiceStatus = original.ServiceStatusDisabled
	ServiceStatusError        ServiceStatus = original.ServiceStatusError
	ServiceStatusProvisioning ServiceStatus = original.ServiceStatusProvisioning
	ServiceStatusRunning      ServiceStatus = original.ServiceStatusRunning
)

type SharedPrivateLinkResourceAsyncOperationResult = original.SharedPrivateLinkResourceAsyncOperationResult

const (
	SharedPrivateLinkResourceAsyncOperationResultFailed    SharedPrivateLinkResourceAsyncOperationResult = original.SharedPrivateLinkResourceAsyncOperationResultFailed
	SharedPrivateLinkResourceAsyncOperationResultRunning   SharedPrivateLinkResourceAsyncOperationResult = original.SharedPrivateLinkResourceAsyncOperationResultRunning
	SharedPrivateLinkResourceAsyncOperationResultSucceeded SharedPrivateLinkResourceAsyncOperationResult = original.SharedPrivateLinkResourceAsyncOperationResultSucceeded
)

type SharedPrivateLinkResourceProvisioningState = original.SharedPrivateLinkResourceProvisioningState

const (
	SharedPrivateLinkResourceProvisioningStateDeleting   SharedPrivateLinkResourceProvisioningState = original.SharedPrivateLinkResourceProvisioningStateDeleting
	SharedPrivateLinkResourceProvisioningStateFailed     SharedPrivateLinkResourceProvisioningState = original.SharedPrivateLinkResourceProvisioningStateFailed
	SharedPrivateLinkResourceProvisioningStateIncomplete SharedPrivateLinkResourceProvisioningState = original.SharedPrivateLinkResourceProvisioningStateIncomplete
	SharedPrivateLinkResourceProvisioningStateSucceeded  SharedPrivateLinkResourceProvisioningState = original.SharedPrivateLinkResourceProvisioningStateSucceeded
	SharedPrivateLinkResourceProvisioningStateUpdating   SharedPrivateLinkResourceProvisioningState = original.SharedPrivateLinkResourceProvisioningStateUpdating
)

type SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatus

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusApproved
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusDisconnected
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusPending
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusRejected
)

type SkuName = original.SkuName

const (
	Basic              SkuName = original.Basic
	Free               SkuName = original.Free
	Standard           SkuName = original.Standard
	Standard2          SkuName = original.Standard2
	Standard3          SkuName = original.Standard3
	StorageOptimizedL1 SkuName = original.StorageOptimizedL1
	StorageOptimizedL2 SkuName = original.StorageOptimizedL2
)

type UnavailableNameReason = original.UnavailableNameReason

const (
	AlreadyExists UnavailableNameReason = original.AlreadyExists
	Invalid       UnavailableNameReason = original.Invalid
)

type AdminKeyResult = original.AdminKeyResult
type AdminKeysClient = original.AdminKeysClient
type AsyncOperationResult = original.AsyncOperationResult
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type IPRule = original.IPRule
type Identity = original.Identity
type ListQueryKeysResult = original.ListQueryKeysResult
type ListQueryKeysResultIterator = original.ListQueryKeysResultIterator
type ListQueryKeysResultPage = original.ListQueryKeysResultPage
type NetworkRuleSet = original.NetworkRuleSet
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionListResultIterator = original.PrivateEndpointConnectionListResultIterator
type PrivateEndpointConnectionListResultPage = original.PrivateEndpointConnectionListResultPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionPropertiesPrivateEndpoint = original.PrivateEndpointConnectionPropertiesPrivateEndpoint
type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState = original.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkResourcesResult = original.PrivateLinkResourcesResult
type ProxyResource = original.ProxyResource
type QueryKey = original.QueryKey
type QueryKeysClient = original.QueryKeysClient
type Resource = original.Resource
type Service = original.Service
type ServiceListResult = original.ServiceListResult
type ServiceListResultIterator = original.ServiceListResultIterator
type ServiceListResultPage = original.ServiceListResultPage
type ServiceProperties = original.ServiceProperties
type ServiceUpdate = original.ServiceUpdate
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ShareablePrivateLinkResourceProperties = original.ShareablePrivateLinkResourceProperties
type ShareablePrivateLinkResourceType = original.ShareablePrivateLinkResourceType
type SharedPrivateLinkResource = original.SharedPrivateLinkResource
type SharedPrivateLinkResourceListResult = original.SharedPrivateLinkResourceListResult
type SharedPrivateLinkResourceListResultIterator = original.SharedPrivateLinkResourceListResultIterator
type SharedPrivateLinkResourceListResultPage = original.SharedPrivateLinkResourceListResultPage
type SharedPrivateLinkResourceProperties = original.SharedPrivateLinkResourceProperties
type SharedPrivateLinkResourcesClient = original.SharedPrivateLinkResourcesClient
type SharedPrivateLinkResourcesCreateOrUpdateFuture = original.SharedPrivateLinkResourcesCreateOrUpdateFuture
type SharedPrivateLinkResourcesDeleteFuture = original.SharedPrivateLinkResourcesDeleteFuture
type Sku = original.Sku
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAdminKeysClient(subscriptionID string) AdminKeysClient {
	return original.NewAdminKeysClient(subscriptionID)
}
func NewAdminKeysClientWithBaseURI(baseURI string, subscriptionID string) AdminKeysClient {
	return original.NewAdminKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewListQueryKeysResultIterator(page ListQueryKeysResultPage) ListQueryKeysResultIterator {
	return original.NewListQueryKeysResultIterator(page)
}
func NewListQueryKeysResultPage(cur ListQueryKeysResult, getNextPage func(context.Context, ListQueryKeysResult) (ListQueryKeysResult, error)) ListQueryKeysResultPage {
	return original.NewListQueryKeysResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListResultIterator(page PrivateEndpointConnectionListResultPage) PrivateEndpointConnectionListResultIterator {
	return original.NewPrivateEndpointConnectionListResultIterator(page)
}
func NewPrivateEndpointConnectionListResultPage(cur PrivateEndpointConnectionListResult, getNextPage func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error)) PrivateEndpointConnectionListResultPage {
	return original.NewPrivateEndpointConnectionListResultPage(cur, getNextPage)
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
func NewQueryKeysClient(subscriptionID string) QueryKeysClient {
	return original.NewQueryKeysClient(subscriptionID)
}
func NewQueryKeysClientWithBaseURI(baseURI string, subscriptionID string) QueryKeysClient {
	return original.NewQueryKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceListResultIterator(page ServiceListResultPage) ServiceListResultIterator {
	return original.NewServiceListResultIterator(page)
}
func NewServiceListResultPage(cur ServiceListResult, getNextPage func(context.Context, ServiceListResult) (ServiceListResult, error)) ServiceListResultPage {
	return original.NewServiceListResultPage(cur, getNextPage)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSharedPrivateLinkResourceListResultIterator(page SharedPrivateLinkResourceListResultPage) SharedPrivateLinkResourceListResultIterator {
	return original.NewSharedPrivateLinkResourceListResultIterator(page)
}
func NewSharedPrivateLinkResourceListResultPage(cur SharedPrivateLinkResourceListResult, getNextPage func(context.Context, SharedPrivateLinkResourceListResult) (SharedPrivateLinkResourceListResult, error)) SharedPrivateLinkResourceListResultPage {
	return original.NewSharedPrivateLinkResourceListResultPage(cur, getNextPage)
}
func NewSharedPrivateLinkResourcesClient(subscriptionID string) SharedPrivateLinkResourcesClient {
	return original.NewSharedPrivateLinkResourcesClient(subscriptionID)
}
func NewSharedPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) SharedPrivateLinkResourcesClient {
	return original.NewSharedPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAdminKeyKindValues() []AdminKeyKind {
	return original.PossibleAdminKeyKindValues()
}
func PossibleHostingModeValues() []HostingMode {
	return original.PossibleHostingModeValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return original.PossiblePrivateLinkServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleServiceStatusValues() []ServiceStatus {
	return original.PossibleServiceStatusValues()
}
func PossibleSharedPrivateLinkResourceAsyncOperationResultValues() []SharedPrivateLinkResourceAsyncOperationResult {
	return original.PossibleSharedPrivateLinkResourceAsyncOperationResultValues()
}
func PossibleSharedPrivateLinkResourceProvisioningStateValues() []SharedPrivateLinkResourceProvisioningState {
	return original.PossibleSharedPrivateLinkResourceProvisioningStateValues()
}
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return original.PossibleSharedPrivateLinkResourceStatusValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleUnavailableNameReasonValues() []UnavailableNameReason {
	return original.PossibleUnavailableNameReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
