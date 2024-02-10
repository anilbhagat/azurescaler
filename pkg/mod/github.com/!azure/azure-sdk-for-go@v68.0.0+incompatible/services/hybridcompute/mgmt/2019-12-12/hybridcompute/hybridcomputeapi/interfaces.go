// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package hybridcomputeapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/hybridcompute/mgmt/2019-12-12/hybridcompute"
	"github.com/Azure/go-autorest/autorest"
)

// MachinesClientAPI contains the set of methods on the MachinesClient type.
type MachinesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters hybridcompute.Machine) (result hybridcompute.Machine, err error)
	Delete(ctx context.Context, resourceGroupName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, name string, expand hybridcompute.InstanceViewTypes) (result hybridcompute.Machine, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result hybridcompute.MachineListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result hybridcompute.MachineListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result hybridcompute.MachineListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result hybridcompute.MachineListResultIterator, err error)
	Reconnect(ctx context.Context, resourceGroupName string, name string, parameters hybridcompute.MachineReconnect) (result hybridcompute.Machine, err error)
	Update(ctx context.Context, resourceGroupName string, name string, parameters hybridcompute.MachineUpdate) (result hybridcompute.Machine, err error)
}

var _ MachinesClientAPI = (*hybridcompute.MachinesClient)(nil)

// MachineExtensionsClientAPI contains the set of methods on the MachineExtensionsClient type.
type MachineExtensionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, extensionName string, extensionParameters hybridcompute.MachineExtension) (result hybridcompute.MachineExtensionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, name string, extensionName string) (result hybridcompute.MachineExtensionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, name string, extensionName string) (result hybridcompute.MachineExtension, err error)
	List(ctx context.Context, resourceGroupName string, name string, expand string) (result hybridcompute.MachineExtensionsListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, name string, expand string) (result hybridcompute.MachineExtensionsListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, name string, extensionName string, extensionParameters hybridcompute.MachineExtensionUpdate) (result hybridcompute.MachineExtensionsUpdateFuture, err error)
}

var _ MachineExtensionsClientAPI = (*hybridcompute.MachineExtensionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result hybridcompute.OperationListResult, err error)
}

var _ OperationsClientAPI = (*hybridcompute.OperationsClient)(nil)