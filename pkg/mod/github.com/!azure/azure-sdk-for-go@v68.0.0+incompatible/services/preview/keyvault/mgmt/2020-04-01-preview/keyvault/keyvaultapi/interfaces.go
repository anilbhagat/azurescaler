// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package keyvaultapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
	"github.com/Azure/go-autorest/autorest"
)

// VaultsClientAPI contains the set of methods on the VaultsClient type.
type VaultsClientAPI interface {
	CheckNameAvailability(ctx context.Context, vaultName keyvault.VaultCheckNameAvailabilityParameters) (result keyvault.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters keyvault.VaultCreateOrUpdateParameters) (result keyvault.VaultsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, vaultName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, vaultName string) (result keyvault.Vault, err error)
	GetDeleted(ctx context.Context, vaultName string, location string) (result keyvault.DeletedVault, err error)
	List(ctx context.Context, top *int32) (result keyvault.ResourceListResultPage, err error)
	ListComplete(ctx context.Context, top *int32) (result keyvault.ResourceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result keyvault.VaultListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result keyvault.VaultListResultIterator, err error)
	ListBySubscription(ctx context.Context, top *int32) (result keyvault.VaultListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, top *int32) (result keyvault.VaultListResultIterator, err error)
	ListDeleted(ctx context.Context) (result keyvault.DeletedVaultListResultPage, err error)
	ListDeletedComplete(ctx context.Context) (result keyvault.DeletedVaultListResultIterator, err error)
	PurgeDeleted(ctx context.Context, vaultName string, location string) (result keyvault.VaultsPurgeDeletedFuture, err error)
	Update(ctx context.Context, resourceGroupName string, vaultName string, parameters keyvault.VaultPatchParameters) (result keyvault.Vault, err error)
	UpdateAccessPolicy(ctx context.Context, resourceGroupName string, vaultName string, operationKind keyvault.AccessPolicyUpdateKind, parameters keyvault.VaultAccessPolicyParameters) (result keyvault.VaultAccessPolicyParameters, err error)
}

var _ VaultsClientAPI = (*keyvault.VaultsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string) (result keyvault.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string) (result keyvault.PrivateEndpointConnection, err error)
	Put(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties keyvault.PrivateEndpointConnection) (result keyvault.PrivateEndpointConnection, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*keyvault.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	ListByVault(ctx context.Context, resourceGroupName string, vaultName string) (result keyvault.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*keyvault.PrivateLinkResourcesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result keyvault.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result keyvault.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*keyvault.OperationsClient)(nil)

// ManagedHsmsClientAPI contains the set of methods on the ManagedHsmsClient type.
type ManagedHsmsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters keyvault.ManagedHsm) (result keyvault.ManagedHsmsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, name string) (result keyvault.ManagedHsmsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, name string) (result keyvault.ManagedHsm, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result keyvault.ManagedHsmListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result keyvault.ManagedHsmListResultIterator, err error)
	ListBySubscription(ctx context.Context, top *int32) (result keyvault.ManagedHsmListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, top *int32) (result keyvault.ManagedHsmListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, name string, parameters keyvault.ManagedHsm) (result keyvault.ManagedHsmsUpdateFuture, err error)
}

var _ ManagedHsmsClientAPI = (*keyvault.ManagedHsmsClient)(nil)
