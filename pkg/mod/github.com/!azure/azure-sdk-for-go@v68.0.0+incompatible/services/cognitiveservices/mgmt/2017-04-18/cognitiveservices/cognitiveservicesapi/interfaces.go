// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package cognitiveservicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/mgmt/2017-04-18/cognitiveservices"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckDomainAvailability(ctx context.Context, parameters cognitiveservices.CheckDomainAvailabilityParameter) (result cognitiveservices.CheckDomainAvailabilityResult, err error)
	CheckSkuAvailability(ctx context.Context, location string, parameters cognitiveservices.CheckSkuAvailabilityParameter) (result cognitiveservices.CheckSkuAvailabilityResultList, err error)
}

var _ BaseClientAPI = (*cognitiveservices.BaseClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, account cognitiveservices.Account) (result cognitiveservices.Account, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	GetProperties(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.Account, err error)
	GetUsages(ctx context.Context, resourceGroupName string, accountName string, filter string) (result cognitiveservices.UsagesResult, err error)
	List(ctx context.Context) (result cognitiveservices.AccountListResultPage, err error)
	ListComplete(ctx context.Context) (result cognitiveservices.AccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result cognitiveservices.AccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result cognitiveservices.AccountListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.AccountKeys, err error)
	ListSkus(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.AccountEnumerateSkusResult, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, parameters cognitiveservices.RegenerateKeyParameters) (result cognitiveservices.AccountKeys, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, account cognitiveservices.Account) (result cognitiveservices.Account, err error)
}

var _ AccountsClientAPI = (*cognitiveservices.AccountsClient)(nil)

// ResourceSkusClientAPI contains the set of methods on the ResourceSkusClient type.
type ResourceSkusClientAPI interface {
	List(ctx context.Context) (result cognitiveservices.ResourceSkusResultPage, err error)
	ListComplete(ctx context.Context) (result cognitiveservices.ResourceSkusResultIterator, err error)
}

var _ ResourceSkusClientAPI = (*cognitiveservices.ResourceSkusClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result cognitiveservices.OperationEntityListResultPage, err error)
	ListComplete(ctx context.Context) (result cognitiveservices.OperationEntityListResultIterator, err error)
}

var _ OperationsClientAPI = (*cognitiveservices.OperationsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, properties cognitiveservices.PrivateEndpointConnection) (result cognitiveservices.PrivateEndpointConnection, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result cognitiveservices.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.PrivateEndpointConnectionListResult, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*cognitiveservices.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, accountName string) (result cognitiveservices.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*cognitiveservices.PrivateLinkResourcesClient)(nil)