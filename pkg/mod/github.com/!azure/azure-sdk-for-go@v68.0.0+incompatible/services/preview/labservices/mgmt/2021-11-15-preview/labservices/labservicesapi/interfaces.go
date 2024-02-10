// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package labservicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/labservices/mgmt/2021-11-15-preview/labservices"
)

// ImagesClientAPI contains the set of methods on the ImagesClient type.
type ImagesClientAPI interface {
	CreateOrUpdate(ctx context.Context, body labservices.Image, resourceGroupName string, labPlanName string, imageName string) (result labservices.Image, err error)
	Get(ctx context.Context, resourceGroupName string, labPlanName string, imageName string) (result labservices.Image, err error)
	ListByLabPlan(ctx context.Context, resourceGroupName string, labPlanName string, filter string) (result labservices.PagedImagesPage, err error)
	ListByLabPlanComplete(ctx context.Context, resourceGroupName string, labPlanName string, filter string) (result labservices.PagedImagesIterator, err error)
	Update(ctx context.Context, body labservices.ImageUpdate, resourceGroupName string, labPlanName string, imageName string) (result labservices.Image, err error)
}

var _ ImagesClientAPI = (*labservices.ImagesClient)(nil)

// LabPlansClientAPI contains the set of methods on the LabPlansClient type.
type LabPlansClientAPI interface {
	CreateOrUpdate(ctx context.Context, body labservices.LabPlan, resourceGroupName string, labPlanName string) (result labservices.LabPlansCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labPlanName string) (result labservices.LabPlansDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labPlanName string) (result labservices.LabPlan, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result labservices.PagedLabPlansPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result labservices.PagedLabPlansIterator, err error)
	ListBySubscription(ctx context.Context, filter string) (result labservices.PagedLabPlansPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string) (result labservices.PagedLabPlansIterator, err error)
	SaveImage(ctx context.Context, body labservices.SaveImageBody, resourceGroupName string, labPlanName string) (result labservices.LabPlansSaveImageFuture, err error)
	Update(ctx context.Context, body labservices.LabPlanUpdate, resourceGroupName string, labPlanName string) (result labservices.LabPlansUpdateFuture, err error)
}

var _ LabPlansClientAPI = (*labservices.LabPlansClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result labservices.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result labservices.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*labservices.OperationsClient)(nil)

// LabsClientAPI contains the set of methods on the LabsClient type.
type LabsClientAPI interface {
	CreateOrUpdate(ctx context.Context, body labservices.Lab, resourceGroupName string, labName string) (result labservices.LabsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string) (result labservices.LabsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string) (result labservices.Lab, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result labservices.PagedLabsPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result labservices.PagedLabsIterator, err error)
	ListBySubscription(ctx context.Context, filter string) (result labservices.PagedLabsPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string) (result labservices.PagedLabsIterator, err error)
	Publish(ctx context.Context, resourceGroupName string, labName string) (result labservices.LabsPublishFuture, err error)
	SyncGroup(ctx context.Context, resourceGroupName string, labName string) (result labservices.LabsSyncGroupFuture, err error)
	Update(ctx context.Context, body labservices.LabUpdate, resourceGroupName string, labName string) (result labservices.LabsUpdateFuture, err error)
}

var _ LabsClientAPI = (*labservices.LabsClient)(nil)

// OperationResultsClientAPI contains the set of methods on the OperationResultsClient type.
type OperationResultsClientAPI interface {
	Get(ctx context.Context, operationResultID string) (result labservices.OperationResult, err error)
}

var _ OperationResultsClientAPI = (*labservices.OperationResultsClient)(nil)

// SchedulesClientAPI contains the set of methods on the SchedulesClient type.
type SchedulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, body labservices.Schedule, resourceGroupName string, labName string, scheduleName string) (result labservices.Schedule, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, scheduleName string) (result labservices.SchedulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, scheduleName string) (result labservices.Schedule, err error)
	ListByLab(ctx context.Context, resourceGroupName string, labName string, filter string) (result labservices.PagedSchedulesPage, err error)
	ListByLabComplete(ctx context.Context, resourceGroupName string, labName string, filter string) (result labservices.PagedSchedulesIterator, err error)
	Update(ctx context.Context, body labservices.ScheduleUpdate, resourceGroupName string, labName string, scheduleName string) (result labservices.Schedule, err error)
}

var _ SchedulesClientAPI = (*labservices.SchedulesClient)(nil)

// UsersClientAPI contains the set of methods on the UsersClient type.
type UsersClientAPI interface {
	CreateOrUpdate(ctx context.Context, body labservices.User, resourceGroupName string, labName string, userName string) (result labservices.UsersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, userName string) (result labservices.UsersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, userName string) (result labservices.User, err error)
	Invite(ctx context.Context, body labservices.InviteBody, resourceGroupName string, labName string, userName string) (result labservices.UsersInviteFuture, err error)
	ListByLab(ctx context.Context, resourceGroupName string, labName string, filter string) (result labservices.PagedUsersPage, err error)
	ListByLabComplete(ctx context.Context, resourceGroupName string, labName string, filter string) (result labservices.PagedUsersIterator, err error)
	Update(ctx context.Context, body labservices.UserUpdate, resourceGroupName string, labName string, userName string) (result labservices.UsersUpdateFuture, err error)
}

var _ UsersClientAPI = (*labservices.UsersClient)(nil)

// VirtualMachinesClientAPI contains the set of methods on the VirtualMachinesClient type.
type VirtualMachinesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string) (result labservices.VirtualMachine, err error)
	ListByLab(ctx context.Context, resourceGroupName string, labName string, filter string) (result labservices.PagedVirtualMachinesPage, err error)
	ListByLabComplete(ctx context.Context, resourceGroupName string, labName string, filter string) (result labservices.PagedVirtualMachinesIterator, err error)
	Redeploy(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string) (result labservices.VirtualMachinesRedeployFuture, err error)
	Reimage(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string) (result labservices.VirtualMachinesReimageFuture, err error)
	ResetPassword(ctx context.Context, body labservices.ResetPasswordBody, resourceGroupName string, labName string, virtualMachineName string) (result labservices.VirtualMachinesResetPasswordFuture, err error)
	Start(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string) (result labservices.VirtualMachinesStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string) (result labservices.VirtualMachinesStopFuture, err error)
}

var _ VirtualMachinesClientAPI = (*labservices.VirtualMachinesClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	ListByLocation(ctx context.Context, location string, filter string) (result labservices.ListUsagesResultPage, err error)
	ListByLocationComplete(ctx context.Context, location string, filter string) (result labservices.ListUsagesResultIterator, err error)
}

var _ UsagesClientAPI = (*labservices.UsagesClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	List(ctx context.Context, filter string) (result labservices.PagedLabServicesSkusPage, err error)
	ListComplete(ctx context.Context, filter string) (result labservices.PagedLabServicesSkusIterator, err error)
}

var _ SkusClientAPI = (*labservices.SkusClient)(nil)
