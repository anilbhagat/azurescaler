//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package subscriptions

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2021-01-01/subscriptions"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type LocationType = original.LocationType

const (
	LocationTypeEdgeZone LocationType = original.LocationTypeEdgeZone
	LocationTypeRegion   LocationType = original.LocationTypeRegion
)

type RegionCategory = original.RegionCategory

const (
	RegionCategoryExtended    RegionCategory = original.RegionCategoryExtended
	RegionCategoryOther       RegionCategory = original.RegionCategoryOther
	RegionCategoryRecommended RegionCategory = original.RegionCategoryRecommended
)

type RegionType = original.RegionType

const (
	RegionTypeLogical  RegionType = original.RegionTypeLogical
	RegionTypePhysical RegionType = original.RegionTypePhysical
)

type ResourceNameStatus = original.ResourceNameStatus

const (
	ResourceNameStatusAllowed  ResourceNameStatus = original.ResourceNameStatusAllowed
	ResourceNameStatusReserved ResourceNameStatus = original.ResourceNameStatusReserved
)

type SpendingLimit = original.SpendingLimit

const (
	SpendingLimitCurrentPeriodOff SpendingLimit = original.SpendingLimitCurrentPeriodOff
	SpendingLimitOff              SpendingLimit = original.SpendingLimitOff
	SpendingLimitOn               SpendingLimit = original.SpendingLimitOn
)

type State = original.State

const (
	StateDeleted  State = original.StateDeleted
	StateDisabled State = original.StateDisabled
	StateEnabled  State = original.StateEnabled
	StatePastDue  State = original.StatePastDue
	StateWarned   State = original.StateWarned
)

type TenantCategory = original.TenantCategory

const (
	TenantCategoryHome        TenantCategory = original.TenantCategoryHome
	TenantCategoryManagedBy   TenantCategory = original.TenantCategoryManagedBy
	TenantCategoryProjectedBy TenantCategory = original.TenantCategoryProjectedBy
)

type BaseClient = original.BaseClient
type CheckResourceNameResult = original.CheckResourceNameResult
type Client = original.Client
type CloudError = original.CloudError
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Location = original.Location
type LocationListResult = original.LocationListResult
type LocationMetadata = original.LocationMetadata
type ManagedByTenant = original.ManagedByTenant
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PairedRegion = original.PairedRegion
type Policies = original.Policies
type ResourceName = original.ResourceName
type Subscription = original.Subscription
type TenantIDDescription = original.TenantIDDescription
type TenantListResult = original.TenantListResult
type TenantListResultIterator = original.TenantListResultIterator
type TenantListResultPage = original.TenantListResultPage
type TenantsClient = original.TenantsClient

func New() BaseClient {
	return original.New()
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewTenantListResultIterator(page TenantListResultPage) TenantListResultIterator {
	return original.NewTenantListResultIterator(page)
}
func NewTenantListResultPage(cur TenantListResult, getNextPage func(context.Context, TenantListResult) (TenantListResult, error)) TenantListResultPage {
	return original.NewTenantListResultPage(cur, getNextPage)
}
func NewTenantsClient() TenantsClient {
	return original.NewTenantsClient()
}
func NewTenantsClientWithBaseURI(baseURI string) TenantsClient {
	return original.NewTenantsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleLocationTypeValues() []LocationType {
	return original.PossibleLocationTypeValues()
}
func PossibleRegionCategoryValues() []RegionCategory {
	return original.PossibleRegionCategoryValues()
}
func PossibleRegionTypeValues() []RegionType {
	return original.PossibleRegionTypeValues()
}
func PossibleResourceNameStatusValues() []ResourceNameStatus {
	return original.PossibleResourceNameStatusValues()
}
func PossibleSpendingLimitValues() []SpendingLimit {
	return original.PossibleSpendingLimitValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleTenantCategoryValues() []TenantCategory {
	return original.PossibleTenantCategoryValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}