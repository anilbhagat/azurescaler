// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package containerserviceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2016-09-30/containerservice"
)

// ContainerServicesClientAPI contains the set of methods on the ContainerServicesClient type.
type ContainerServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters containerservice.ContainerService) (result containerservice.ContainerServicesCreateOrUpdateFutureType, err error)
	Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.ContainerServicesDeleteFutureType, err error)
	Get(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.ContainerService, err error)
	List(ctx context.Context) (result containerservice.ListResultPage, err error)
	ListComplete(ctx context.Context) (result containerservice.ListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.ListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerservice.ListResultIterator, err error)
}

var _ ContainerServicesClientAPI = (*containerservice.ContainerServicesClient)(nil)
