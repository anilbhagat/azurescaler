// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package consumptionapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2018-05-31/consumption"
)

// PriceSheetClientAPI contains the set of methods on the PriceSheetClient type.
type PriceSheetClientAPI interface {
	Get(ctx context.Context, expand string, skiptoken string, top *int32) (result consumption.PriceSheetResult, err error)
	GetByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, skiptoken string, top *int32) (result consumption.PriceSheetResult, err error)
}

var _ PriceSheetClientAPI = (*consumption.PriceSheetClient)(nil)

// UsageDetailsClientAPI contains the set of methods on the UsageDetailsClient type.
type UsageDetailsClientAPI interface {
	List(ctx context.Context, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListComplete(ctx context.Context, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultIterator, err error)
	ListByBillingAccount(ctx context.Context, billingAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultIterator, err error)
	ListByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListByBillingPeriodComplete(ctx context.Context, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultIterator, err error)
	ListByDepartment(ctx context.Context, departmentID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByDepartmentComplete(ctx context.Context, departmentID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultIterator, err error)
	ListByEnrollmentAccount(ctx context.Context, enrollmentAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByEnrollmentAccountComplete(ctx context.Context, enrollmentAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultIterator, err error)
	ListForBillingPeriodByBillingAccount(ctx context.Context, billingAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByBillingAccountComplete(ctx context.Context, billingAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultIterator, err error)
	ListForBillingPeriodByDepartment(ctx context.Context, departmentID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByDepartmentComplete(ctx context.Context, departmentID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultIterator, err error)
	ListForBillingPeriodByEnrollmentAccount(ctx context.Context, enrollmentAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByEnrollmentAccountComplete(ctx context.Context, enrollmentAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultIterator, err error)
}

var _ UsageDetailsClientAPI = (*consumption.UsageDetailsClient)(nil)

// ForecastsClientAPI contains the set of methods on the ForecastsClient type.
type ForecastsClientAPI interface {
	List(ctx context.Context, filter string) (result consumption.ForecastsListResult, err error)
}

var _ ForecastsClientAPI = (*consumption.ForecastsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result consumption.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result consumption.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*consumption.OperationsClient)(nil)
