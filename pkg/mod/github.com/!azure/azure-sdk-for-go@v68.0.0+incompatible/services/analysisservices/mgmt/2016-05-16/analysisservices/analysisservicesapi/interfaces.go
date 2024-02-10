// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package analysisservicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/analysisservices/mgmt/2016-05-16/analysisservices"
	"github.com/Azure/go-autorest/autorest"
)

// ServersClientAPI contains the set of methods on the ServersClient type.
type ServersClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, serverParameters analysisservices.CheckServerNameAvailabilityParameters) (result analysisservices.CheckServerNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, serverName string, serverParameters analysisservices.Server) (result analysisservices.ServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.ServersDeleteFuture, err error)
	GetDetails(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.Server, err error)
	List(ctx context.Context) (result analysisservices.Servers, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result analysisservices.Servers, err error)
	ListOperationResults(ctx context.Context, location string, operationID string) (result autorest.Response, err error)
	ListOperationStatuses(ctx context.Context, location string, operationID string) (result analysisservices.OperationStatus, err error)
	ListSkusForExisting(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.SkuEnumerationForExistingResourceResult, err error)
	ListSkusForNew(ctx context.Context) (result analysisservices.SkuEnumerationForNewResourceResult, err error)
	Resume(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.ServersResumeFuture, err error)
	Suspend(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.ServersSuspendFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, serverUpdateParameters analysisservices.ServerUpdateParameters) (result analysisservices.ServersUpdateFuture, err error)
}

var _ ServersClientAPI = (*analysisservices.ServersClient)(nil)
