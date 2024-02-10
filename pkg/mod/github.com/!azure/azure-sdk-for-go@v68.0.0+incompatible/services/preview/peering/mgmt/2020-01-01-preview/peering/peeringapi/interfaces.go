// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package peeringapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/peering/mgmt/2020-01-01-preview/peering"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckServiceProviderAvailability(ctx context.Context, checkServiceProviderAvailabilityInput peering.CheckServiceProviderAvailabilityInput) (result peering.String, err error)
}

var _ BaseClientAPI = (*peering.BaseClient)(nil)

// LegacyPeeringsClientAPI contains the set of methods on the LegacyPeeringsClient type.
type LegacyPeeringsClientAPI interface {
	List(ctx context.Context, peeringLocation string, kind string, asn *int32) (result peering.ListResultPage, err error)
	ListComplete(ctx context.Context, peeringLocation string, kind string, asn *int32) (result peering.ListResultIterator, err error)
}

var _ LegacyPeeringsClientAPI = (*peering.LegacyPeeringsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result peering.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result peering.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*peering.OperationsClient)(nil)

// PeerAsnsClientAPI contains the set of methods on the PeerAsnsClient type.
type PeerAsnsClientAPI interface {
	CreateOrUpdate(ctx context.Context, peerAsnName string, peerAsn peering.PeerAsn) (result peering.PeerAsn, err error)
	Delete(ctx context.Context, peerAsnName string) (result autorest.Response, err error)
	Get(ctx context.Context, peerAsnName string) (result peering.PeerAsn, err error)
	ListBySubscription(ctx context.Context) (result peering.PeerAsnListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result peering.PeerAsnListResultIterator, err error)
}

var _ PeerAsnsClientAPI = (*peering.PeerAsnsClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	List(ctx context.Context, kind string, directPeeringType string) (result peering.LocationListResultPage, err error)
	ListComplete(ctx context.Context, kind string, directPeeringType string) (result peering.LocationListResultIterator, err error)
}

var _ LocationsClientAPI = (*peering.LocationsClient)(nil)

// RegisteredAsnsClientAPI contains the set of methods on the RegisteredAsnsClient type.
type RegisteredAsnsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string, registeredAsn peering.RegisteredAsn) (result peering.RegisteredAsn, err error)
	Delete(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, peeringName string, registeredAsnName string) (result peering.RegisteredAsn, err error)
	ListByPeering(ctx context.Context, resourceGroupName string, peeringName string) (result peering.RegisteredAsnListResultPage, err error)
	ListByPeeringComplete(ctx context.Context, resourceGroupName string, peeringName string) (result peering.RegisteredAsnListResultIterator, err error)
}

var _ RegisteredAsnsClientAPI = (*peering.RegisteredAsnsClient)(nil)

// RegisteredPrefixesClientAPI contains the set of methods on the RegisteredPrefixesClient type.
type RegisteredPrefixesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, registeredPrefix peering.RegisteredPrefix) (result peering.RegisteredPrefix, err error)
	Delete(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string) (result peering.RegisteredPrefix, err error)
	ListByPeering(ctx context.Context, resourceGroupName string, peeringName string) (result peering.RegisteredPrefixListResultPage, err error)
	ListByPeeringComplete(ctx context.Context, resourceGroupName string, peeringName string) (result peering.RegisteredPrefixListResultIterator, err error)
}

var _ RegisteredPrefixesClientAPI = (*peering.RegisteredPrefixesClient)(nil)

// PeeringsClientAPI contains the set of methods on the PeeringsClient type.
type PeeringsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, peering peering.Model) (result peering.Model, err error)
	Delete(ctx context.Context, resourceGroupName string, peeringName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, peeringName string) (result peering.Model, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result peering.ListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result peering.ListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result peering.ListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result peering.ListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, peeringName string, tags peering.ResourceTags) (result peering.Model, err error)
}

var _ PeeringsClientAPI = (*peering.PeeringsClient)(nil)

// ServiceCountriesClientAPI contains the set of methods on the ServiceCountriesClient type.
type ServiceCountriesClientAPI interface {
	List(ctx context.Context) (result peering.ServiceCountryListResultPage, err error)
	ListComplete(ctx context.Context) (result peering.ServiceCountryListResultIterator, err error)
}

var _ ServiceCountriesClientAPI = (*peering.ServiceCountriesClient)(nil)

// ServiceLocationsClientAPI contains the set of methods on the ServiceLocationsClient type.
type ServiceLocationsClientAPI interface {
	List(ctx context.Context, country string) (result peering.ServiceLocationListResultPage, err error)
	ListComplete(ctx context.Context, country string) (result peering.ServiceLocationListResultIterator, err error)
}

var _ ServiceLocationsClientAPI = (*peering.ServiceLocationsClient)(nil)

// PrefixesClientAPI contains the set of methods on the PrefixesClient type.
type PrefixesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, peeringServicePrefix peering.ServicePrefix) (result peering.ServicePrefix, err error)
	Delete(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, expand string) (result peering.ServicePrefix, err error)
	ListByPeeringService(ctx context.Context, resourceGroupName string, peeringServiceName string, expand string) (result peering.ServicePrefixListResultPage, err error)
	ListByPeeringServiceComplete(ctx context.Context, resourceGroupName string, peeringServiceName string, expand string) (result peering.ServicePrefixListResultIterator, err error)
}

var _ PrefixesClientAPI = (*peering.PrefixesClient)(nil)

// ServiceProvidersClientAPI contains the set of methods on the ServiceProvidersClient type.
type ServiceProvidersClientAPI interface {
	List(ctx context.Context) (result peering.ServiceProviderListResultPage, err error)
	ListComplete(ctx context.Context) (result peering.ServiceProviderListResultIterator, err error)
}

var _ ServiceProvidersClientAPI = (*peering.ServiceProvidersClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringServiceName string, peeringService peering.Service) (result peering.Service, err error)
	Delete(ctx context.Context, resourceGroupName string, peeringServiceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, peeringServiceName string) (result peering.Service, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result peering.ServiceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result peering.ServiceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result peering.ServiceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result peering.ServiceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, peeringServiceName string, tags peering.ResourceTags) (result peering.Service, err error)
}

var _ ServicesClientAPI = (*peering.ServicesClient)(nil)
