// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package mediaapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/mediaservices/mgmt/2015-10-01/media"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result media.OperationListResult, err error)
}

var _ OperationsClientAPI = (*media.OperationsClient)(nil)

// ServiceClientAPI contains the set of methods on the ServiceClient type.
type ServiceClientAPI interface {
	CheckNameAvailability(ctx context.Context, parameters media.CheckNameAvailabilityInput) (result media.CheckNameAvailabilityOutput, err error)
	Create(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters media.Service) (result media.Service, err error)
	Delete(ctx context.Context, resourceGroupName string, mediaServiceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, mediaServiceName string) (result media.Service, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result media.ServiceCollection, err error)
	ListKeys(ctx context.Context, resourceGroupName string, mediaServiceName string) (result media.ServiceKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters media.RegenerateKeyInput) (result media.RegenerateKeyOutput, err error)
	SyncStorageKeys(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters media.SyncStorageKeysInput) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters media.Service) (result media.Service, err error)
}

var _ ServiceClientAPI = (*media.ServiceClient)(nil)