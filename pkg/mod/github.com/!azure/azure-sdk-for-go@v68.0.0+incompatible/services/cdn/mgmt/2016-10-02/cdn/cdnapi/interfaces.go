// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package cdnapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2016-10-02/cdn"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput cdn.CheckNameAvailabilityInput) (result cdn.CheckNameAvailabilityOutput, err error)
	ListOperations(ctx context.Context) (result cdn.OperationListResultPage, err error)
	ListOperationsComplete(ctx context.Context) (result cdn.OperationListResultIterator, err error)
	ListResourceUsage(ctx context.Context) (result cdn.ResourceUsageListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context) (result cdn.ResourceUsageListResultIterator, err error)
}

var _ BaseClientAPI = (*cdn.BaseClient)(nil)

// ProfilesClientAPI contains the set of methods on the ProfilesClient type.
type ProfilesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, profile cdn.Profile) (result cdn.ProfilesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string) (result cdn.ProfilesDeleteFuture, err error)
	GenerateSsoURI(ctx context.Context, resourceGroupName string, profileName string) (result cdn.SsoURI, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string) (result cdn.Profile, err error)
	GetSupportedOptimizationTypes(ctx context.Context, resourceGroupName string, profileName string) (result cdn.SupportedOptimizationTypesResult, err error)
	List(ctx context.Context) (result cdn.ProfileListResultPage, err error)
	ListComplete(ctx context.Context) (result cdn.ProfileListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result cdn.ProfileListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result cdn.ProfileListResultIterator, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string) (result cdn.ResourceUsageListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string) (result cdn.ResourceUsageListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, profileUpdateParameters cdn.ProfileUpdateParameters) (result cdn.ProfilesUpdateFuture, err error)
}

var _ ProfilesClientAPI = (*cdn.ProfilesClient)(nil)

// EndpointsClientAPI contains the set of methods on the EndpointsClient type.
type EndpointsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint cdn.Endpoint) (result cdn.EndpointsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.Endpoint, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.EndpointListResultPage, err error)
	ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result cdn.EndpointListResultIterator, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.ResourceUsageListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.ResourceUsageListResultIterator, err error)
	LoadContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths cdn.LoadParameters) (result cdn.EndpointsLoadContentFuture, err error)
	PurgeContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths cdn.PurgeParameters) (result cdn.EndpointsPurgeContentFuture, err error)
	Start(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties cdn.EndpointUpdateParameters) (result cdn.EndpointsUpdateFuture, err error)
	ValidateCustomDomain(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties cdn.ValidateCustomDomainInput) (result cdn.ValidateCustomDomainOutput, err error)
}

var _ EndpointsClientAPI = (*cdn.EndpointsClient)(nil)

// OriginsClientAPI contains the set of methods on the OriginsClient type.
type OriginsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string) (result cdn.Origin, err error)
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginListResultPage, err error)
	ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties cdn.OriginUpdateParameters) (result cdn.OriginsUpdateFuture, err error)
}

var _ OriginsClientAPI = (*cdn.OriginsClient)(nil)

// CustomDomainsClientAPI contains the set of methods on the CustomDomainsClient type.
type CustomDomainsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties cdn.CustomDomainParameters) (result cdn.CustomDomainsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomainsDeleteFuture, err error)
	DisableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomain, err error)
	EnableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomain, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomain, err error)
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.CustomDomainListResultPage, err error)
	ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.CustomDomainListResultIterator, err error)
}

var _ CustomDomainsClientAPI = (*cdn.CustomDomainsClient)(nil)

// EdgeNodesClientAPI contains the set of methods on the EdgeNodesClient type.
type EdgeNodesClientAPI interface {
	List(ctx context.Context) (result cdn.EdgenodeResultPage, err error)
	ListComplete(ctx context.Context) (result cdn.EdgenodeResultIterator, err error)
}

var _ EdgeNodesClientAPI = (*cdn.EdgeNodesClient)(nil)
