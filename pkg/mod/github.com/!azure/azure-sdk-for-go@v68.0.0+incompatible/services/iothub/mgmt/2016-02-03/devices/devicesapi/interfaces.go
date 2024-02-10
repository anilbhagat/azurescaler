// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package devicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2016-02-03/devices"
	"github.com/Azure/go-autorest/autorest"
)

// IotHubResourceClientAPI contains the set of methods on the IotHubResourceClient type.
type IotHubResourceClientAPI interface {
	CheckNameAvailability(ctx context.Context, operationInputs devices.OperationInputs) (result devices.IotHubNameAvailabilityInfo, err error)
	CreateEventHubConsumerGroup(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string, name string) (result devices.EventHubConsumerGroupInfo, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, iotHubDescription devices.IotHubDescription) (result devices.IotHubResourceCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubResourceDeleteFuture, err error)
	DeleteEventHubConsumerGroup(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string, name string) (result autorest.Response, err error)
	ExportDevices(ctx context.Context, resourceGroupName string, resourceName string, exportDevicesParameters devices.ExportDevicesRequest) (result devices.JobResponse, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubDescription, err error)
	GetEventHubConsumerGroup(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string, name string) (result devices.EventHubConsumerGroupInfo, err error)
	GetJob(ctx context.Context, resourceGroupName string, resourceName string, jobID string) (result devices.JobResponse, err error)
	GetKeysForKeyName(ctx context.Context, resourceGroupName string, resourceName string, keyName string) (result devices.SharedAccessSignatureAuthorizationRule, err error)
	GetQuotaMetrics(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubQuotaMetricInfoListResultPage, err error)
	GetQuotaMetricsComplete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubQuotaMetricInfoListResultIterator, err error)
	GetStats(ctx context.Context, resourceGroupName string, resourceName string) (result devices.RegistryStatistics, err error)
	GetValidSkus(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubSkuDescriptionListResultPage, err error)
	GetValidSkusComplete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubSkuDescriptionListResultIterator, err error)
	ImportDevices(ctx context.Context, resourceGroupName string, resourceName string, importDevicesParameters devices.ImportDevicesRequest) (result devices.JobResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result devices.IotHubDescriptionListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result devices.IotHubDescriptionListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result devices.IotHubDescriptionListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result devices.IotHubDescriptionListResultIterator, err error)
	ListEventHubConsumerGroups(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string) (result devices.EventHubConsumerGroupsListResultPage, err error)
	ListEventHubConsumerGroupsComplete(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string) (result devices.EventHubConsumerGroupsListResultIterator, err error)
	ListJobs(ctx context.Context, resourceGroupName string, resourceName string) (result devices.JobResponseListResultPage, err error)
	ListJobsComplete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.JobResponseListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, resourceName string) (result devices.SharedAccessSignatureAuthorizationRuleListResultPage, err error)
	ListKeysComplete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.SharedAccessSignatureAuthorizationRuleListResultIterator, err error)
}

var _ IotHubResourceClientAPI = (*devices.IotHubResourceClient)(nil)
