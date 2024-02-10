// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package recoveryservicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2021-08-01/recoveryservices"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	GetOperationResult(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (result recoveryservices.Vault, err error)
	GetOperationStatus(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (result recoveryservices.OperationResource, err error)
}

var _ BaseClientAPI = (*recoveryservices.BaseClient)(nil)

// VaultCertificatesClientAPI contains the set of methods on the VaultCertificatesClient type.
type VaultCertificatesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, vaultName string, certificateName string, certificateRequest recoveryservices.CertificateRequest) (result recoveryservices.VaultCertificateResponse, err error)
}

var _ VaultCertificatesClientAPI = (*recoveryservices.VaultCertificatesClient)(nil)

// RegisteredIdentitiesClientAPI contains the set of methods on the RegisteredIdentitiesClient type.
type RegisteredIdentitiesClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, vaultName string, identityName string) (result autorest.Response, err error)
}

var _ RegisteredIdentitiesClientAPI = (*recoveryservices.RegisteredIdentitiesClient)(nil)

// ReplicationUsagesClientAPI contains the set of methods on the ReplicationUsagesClient type.
type ReplicationUsagesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.ReplicationUsageList, err error)
}

var _ ReplicationUsagesClientAPI = (*recoveryservices.ReplicationUsagesClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, vaultName string, privateLinkResourceName string) (result recoveryservices.PrivateLinkResource, err error)
	List(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.PrivateLinkResourcesPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.PrivateLinkResourcesIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*recoveryservices.PrivateLinkResourcesClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckNameAvailability(ctx context.Context, resourceGroupName string, location string, input recoveryservices.CheckNameAvailabilityParameters) (result recoveryservices.CheckNameAvailabilityResult, err error)
}

var _ ClientAPI = (*recoveryservices.Client)(nil)

// VaultsClientAPI contains the set of methods on the VaultsClient type.
type VaultsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, vault recoveryservices.Vault) (result recoveryservices.VaultsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, vaultName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.Vault, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result recoveryservices.VaultListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result recoveryservices.VaultListIterator, err error)
	ListBySubscriptionID(ctx context.Context) (result recoveryservices.VaultListPage, err error)
	ListBySubscriptionIDComplete(ctx context.Context) (result recoveryservices.VaultListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, vaultName string, vault recoveryservices.PatchVault) (result recoveryservices.VaultsUpdateFuture, err error)
}

var _ VaultsClientAPI = (*recoveryservices.VaultsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result recoveryservices.ClientDiscoveryResponsePage, err error)
	ListComplete(ctx context.Context) (result recoveryservices.ClientDiscoveryResponseIterator, err error)
}

var _ OperationsClientAPI = (*recoveryservices.OperationsClient)(nil)

// VaultExtendedInfoClientAPI contains the set of methods on the VaultExtendedInfoClient type.
type VaultExtendedInfoClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, resourceResourceExtendedInfoDetails recoveryservices.VaultExtendedInfoResource) (result recoveryservices.VaultExtendedInfoResource, err error)
	Get(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.VaultExtendedInfoResource, err error)
	Update(ctx context.Context, resourceGroupName string, vaultName string, resourceResourceExtendedInfoDetails recoveryservices.VaultExtendedInfoResource) (result recoveryservices.VaultExtendedInfoResource, err error)
}

var _ VaultExtendedInfoClientAPI = (*recoveryservices.VaultExtendedInfoClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	ListByVaults(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.VaultUsageList, err error)
}

var _ UsagesClientAPI = (*recoveryservices.UsagesClient)(nil)